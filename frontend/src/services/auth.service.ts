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
      if (response.token) {
        localStorage.setItem('token', response.token);
        localStorage.setItem('user', JSON.stringify(response.user));
      }
      
      return response;
    } catch (error) {
      console.error('AuthService login error:', error);
      throw error;
    }
  },

  async signup(data: SignupRequest): Promise<AuthResponse> {
    // Using the authApi.register function directly
    const response = await authApi.register(data);
    if (response.token) {
      localStorage.setItem('token', response.token);
      localStorage.setItem('user', JSON.stringify(response.user));
    }
    return response;
  },

  async logout(): Promise<void> {
    localStorage.removeItem('token');
    localStorage.removeItem('user');
  },

  getCurrentUser() {
    const userStr = localStorage.getItem('user');
    return userStr ? JSON.parse(userStr) : null;
  },

  getToken(): string | null {
    return localStorage.getItem('token');
  },

  isAuthenticated(): boolean {
    return !!this.getToken();
  },
};