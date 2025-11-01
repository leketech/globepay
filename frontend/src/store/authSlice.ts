import { createSlice, createAsyncThunk, PayloadAction } from '@reduxjs/toolkit';
import { authService, LoginRequest, SignupRequest, AuthResponse } from '../services/auth.service';
import { AuthState, User } from '../types';

// Create a function to map the API user response to our User type
const mapApiUserToUser = (apiUser: any): User => {
  console.log('Mapping API user to frontend user:', apiUser);
  
  const user: User = {
    id: apiUser.id,
    email: apiUser.email,
    firstName: apiUser.first_name || apiUser.firstName || '',
    lastName: apiUser.last_name || apiUser.lastName || '',
    phoneNumber: apiUser.phone_number || apiUser.phoneNumber || '',
    dateOfBirth: apiUser.date_of_birth || apiUser.dateOfBirth || '',
    country: apiUser.country || '',
    kycStatus: apiUser.kyc_status || apiUser.kycStatus || '',
    accountStatus: apiUser.account_status || apiUser.accountStatus || '',
    createdAt: apiUser.created_at || apiUser.createdAt || '',
    updatedAt: apiUser.updated_at || apiUser.updatedAt || '',
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
    } catch (error: any) {
      console.error('Login error in authService:', error);
      // Try to extract more detailed error information
      let errorMessage = 'Login failed';
      if (error.response?.data?.message) {
        errorMessage = error.response.data.message;
      } else if (error.response?.data?.error) {
        errorMessage = error.response.data.error;
      } else if (error.message) {
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
      const response = await authService.signup(userData);
      return response;
    } catch (error: any) {
      return rejectWithValue(error.response?.data?.message || 'Signup failed');
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