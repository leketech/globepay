import { createSlice, createAsyncThunk, PayloadAction } from '@reduxjs/toolkit';
import { authService, LoginRequest, SignupRequest, AuthResponse } from '../services/auth.service';
import { AuthState, User } from '../types';

interface ApiUser {
  id: string;
  email: string;
  firstName: string;
  lastName: string;
  phoneNumber?: string;
  dateOfBirth?: string;
  country?: string;
  kycStatus?: string;
  accountStatus?: string;
  createdAt?: string;
  updatedAt?: string;
}

// Create a function to map the API user response to our User type
const mapApiUserToUser = (apiUser: ApiUser): User => {
  console.log('Mapping API user to frontend user:', apiUser);

  const user: User = {
    id: apiUser.id,
    email: apiUser.email,
    firstName: apiUser.firstName || '',
    lastName: apiUser.lastName || '',
    phoneNumber: apiUser.phoneNumber || '',
    dateOfBirth: apiUser.dateOfBirth || '',
    country: apiUser.country || '',
    kycStatus: apiUser.kycStatus || '',
    accountStatus: apiUser.accountStatus || '',
    createdAt: apiUser.createdAt || '',
    updatedAt: apiUser.updatedAt || '',
  };

  console.log('Mapped user:', user);
  return user;
};

const initialState: AuthState = {
  user: authService.getCurrentUser(),
  token: authService.getToken(),
  accounts: [],
  isAuthenticated: authService.isAuthenticated(),
  loading: false,
  error: null,
};

export const login = createAsyncThunk(
  'auth/login',
  async (credentials: LoginRequest, { rejectWithValue }) => {
    try {
      console.log('AuthService login called with:', credentials);
      const response = await authService.login(credentials);
      console.log('AuthService login response:', response);
      return response;
    } catch (error: unknown) {
      console.error('Login error in authService:', error);
      // Try to extract more detailed error information
      let errorMessage = 'Login failed';
      if (error instanceof Error) {
        errorMessage = error.message;
      }
      return rejectWithValue(errorMessage);
    }
  }
);

export const signup = createAsyncThunk(
  'auth/signup',
  async (userData: SignupRequest, { rejectWithValue }) => {
    try {
      console.log('authSlice signup called with userData:', userData);
      const response = await authService.signup(userData);
      console.log('authSlice signup response:', response);
      return response;
    } catch (error: unknown) {
      console.error('Signup error in authSlice:', error);
      if (error instanceof Error) {
        return rejectWithValue(error.message || 'Signup failed');
      }
      return rejectWithValue('Signup failed');
    }
  }
);

export const logout = createAsyncThunk('auth/logout', async () => {
  await authService.logout();
});

const authSlice = createSlice({
  name: 'auth',
  initialState,
  reducers: {
    clearError: (state) => {
      state.error = null;
    },
  },
  extraReducers: (builder) => {
    builder
      // Login
      .addCase(login.pending, (state) => {
        state.loading = true;
        state.error = null;
      })
      .addCase(login.fulfilled, (state, action: PayloadAction<AuthResponse>) => {
        state.loading = false;
        state.isAuthenticated = true;
        state.user = mapApiUserToUser(action.payload.user);
        state.token = action.payload.token;
      })
      .addCase(login.rejected, (state, action) => {
        state.loading = false;
        state.error = action.payload as string;
      })
      // Signup
      .addCase(signup.pending, (state) => {
        state.loading = true;
        state.error = null;
      })
      .addCase(signup.fulfilled, (state, action: PayloadAction<AuthResponse>) => {
        state.loading = false;
        state.isAuthenticated = true;
        state.user = mapApiUserToUser(action.payload.user);
        state.token = action.payload.token;
      })
      .addCase(signup.rejected, (state, action) => {
        state.loading = false;
        state.error = action.payload as string;
      })
      // Logout
      .addCase(logout.fulfilled, (state) => {
        state.user = null;
        state.token = null;
        state.accounts = [];
        state.isAuthenticated = false;
      });
  },
});

export const { clearError } = authSlice.actions;
export default authSlice.reducer;
