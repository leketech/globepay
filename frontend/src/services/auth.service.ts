import { authApi } from './api';

export interface LoginRequest {
  email: string;
  password: string;
}

export interface SignupRequest {
  email: string;
  password: string;
  firstName: string;
  lastName: string;
  phoneNumber?: string;
  dateOfBirth?: string;
  country?: string;
}

export interface AuthResponse {
  token: string;
  user: {
    id: string;
    email: string;
    firstName: string;
    lastName: string;
  };
}

// Helper function to safely access localStorage
const getStorage = () => {
  if (typeof window === 'undefined' || typeof localStorage === 'undefined') {
    return null;
  }
  return localStorage;
};

export const authService = {
  async login(data: LoginRequest): Promise<AuthResponse> {
    console.log('AuthService login called with data:', data);
    try {
      // Using the authApi.login function directly
      const response = await authApi.login(data.email, data.password);
      console.log('AuthService login response:', response);

      // Ensure the response has the expected structure
      if (!response || !response.token || !response.user) {
        throw new Error('Invalid login response');
      }

      // Store in localStorage
      const storage = getStorage();
      if (storage && response.token) {
        storage.setItem('token', response.token);
        storage.setItem('user', JSON.stringify(response.user));
      }

      return response;
    } catch (error) {
      console.error('AuthService login error:', error);
      throw error;
    }
  },

  async signup(data: SignupRequest): Promise<AuthResponse> {
    console.log('AuthService signup called with data:', data);
    // Using the authApi.register function directly
    const response = await authApi.register(data);
    console.log('AuthService signup response:', response);
    const storage = getStorage();
    if (storage && response.token) {
      storage.setItem('token', response.token);
      storage.setItem('user', JSON.stringify(response.user));
    }
    return response;
  },

  async logout(): Promise<void> {
    const storage = getStorage();
    if (storage) {
      storage.removeItem('token');
      storage.removeItem('user');
    }
  },

  getCurrentUser() {
    const storage = getStorage();
    if (!storage) return null;
    const userStr = storage.getItem('user');
    return userStr ? JSON.parse(userStr) : null;
  },

  getToken(): string | null {
    const storage = getStorage();
    if (!storage) return null;
    return storage.getItem('token');
  },

  isAuthenticated(): boolean {
    return !!this.getToken();
  },
};
