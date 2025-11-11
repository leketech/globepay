// Simple API service for Globepay
import { Transfer } from '../types';

// Use a default API base URL for testing environments
const API_BASE_URL =
  typeof process !== 'undefined' && process.env.VITE_API_URL
    ? process.env.VITE_API_URL.replace(/\/$/, '')
    : typeof window !== 'undefined' && window.location
      ? '/api'
      : 'http://localhost:8080';

// Helper function for API requests
const apiRequest = async (endpoint: string, options: RequestInit = {}) => {
  const isHealthCheck = endpoint.startsWith('/health');
  let url = `${API_BASE_URL}${endpoint}`;

  // Handle health check URLs in development
  if (isHealthCheck && typeof process !== 'undefined' && process.env.NODE_ENV === 'development') {
    url = `http://localhost:8080${endpoint}`;
  }

  const config: RequestInit = {
    headers: {
      'Content-Type': 'application/json',
      ...options.headers,
    },
    ...options,
  };

  try {
    console.log(`Making API request to: ${url}`, config);
    const response = await fetch(url, config);

    console.log(`API response status: ${response.status}`);

    const text = await response.text();
    const data = text ? JSON.parse(text) : {};

    if (!response.ok) {
      const errorData = data;
      console.error('API request failed:', {
        status: response.status,
        statusText: response.statusText,
        errorData,
      });
      throw {
        status: response.status,
        message: errorData?.error || response.statusText,
        details: errorData,
      };
    }

    console.log('API response data:', data);
    return data;
  } catch (error) {
    console.error('API request failed:', error);
    throw error;
  }
};

// Helper function to filter out empty optional fields
const filterEmptyFields = <T extends object>(obj: T): Partial<T> => {
  const filtered: Partial<T> = {};
  for (const key in obj) {
    if (obj[key] !== '' && obj[key] !== null && obj[key] !== undefined) {
      filtered[key] = obj[key];
    }
  }
  console.log('Filtered data:', filtered);
  return filtered;
};

interface RegisterUserData {
  email: string;
  password: string;
  firstName: string;
  lastName: string;
  phoneNumber?: string;
  dateOfBirth?: string;
  country?: string;
}

interface UpdateUserData {
  firstName?: string;
  lastName?: string;
  phoneNumber?: string;
  dateOfBirth?: string;
  country?: string;
}

interface UserPreferences {
  language?: string;
  timezone?: string;
  notifications?: {
    email: boolean;
    sms: boolean;
  };
}

// Auth API
export const authApi = {
  login: async (email: string, password: string) => {
    // Ensure we're sending properly formatted JSON
    const body = JSON.stringify({ email, password });
    console.log('Login request body:', body);
    return apiRequest('/v1/auth/login', {
      method: 'POST',
      body: body,
    });
  },

  register: async (userData: RegisterUserData) => {
    // Filter out empty optional fields
    const filteredData = filterEmptyFields(userData);
    // Ensure we're sending properly formatted JSON with correct field names
    const body = JSON.stringify({
      email: filteredData.email,
      password: filteredData.password,
      firstName: filteredData.firstName,
      lastName: filteredData.lastName,
      phoneNumber: filteredData.phoneNumber,
      dateOfBirth: filteredData.dateOfBirth,
      country: filteredData.country,
    });
    console.log('Register request body:', body);
    return apiRequest('/v1/auth/register', {
      method: 'POST',
      body: body,
    });
  },

  refreshToken: async (refreshToken: string) => {
    // Ensure we're sending properly formatted JSON
    const body = JSON.stringify({ refreshToken });
    return apiRequest('/v1/auth/refresh', {
      method: 'POST',
      body: body,
    });
  },
};

// User API
export const userApi = {
  getProfile: async (token: string) => {
    return apiRequest('/v1/user/profile', {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
  },

  updateProfile: async (token: string, userData: UpdateUserData) => {
    return apiRequest('/v1/user/profile', {
      method: 'PUT',
      headers: {
        Authorization: `Bearer ${token}`,
      },
      body: JSON.stringify(userData),
    });
  },

  getAccounts: async (token: string) => {
    return apiRequest('/v1/user/accounts', {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
  },

  createAccount: async (token: string, currency: string) => {
    return apiRequest('/v1/user/accounts', {
      method: 'POST',
      headers: {
        Authorization: `Bearer ${token}`,
      },
      body: JSON.stringify({ currency }),
    });
  },

  getUserPreferences: async (token: string) => {
    return apiRequest('/v1/user/preferences', {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
  },

  updateUserPreferences: async (token: string, preferences: Partial<UserPreferences>) => {
    return apiRequest('/v1/user/preferences', {
      method: 'PUT',
      headers: {
        Authorization: `Bearer ${token}`,
      },
      body: JSON.stringify(preferences),
    });
  },
};

// Transfer API
export const transferApi = {
  getTransfers: async (token: string, page = 1, limit = 10) => {
    return apiRequest(`/v1/transfers?page=${page}&limit=${limit}`, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
  },

  getTransfer: async (token: string, transferId: string) => {
    return apiRequest(`/v1/transfers/${transferId}`, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
  },

  createTransfer: async (token: string, transferData: Transfer) => {
    return apiRequest('/v1/transfers', {
      method: 'POST',
      headers: {
        Authorization: `Bearer ${token}`,
      },
      body: JSON.stringify(transferData),
    });
  },

  cancelTransfer: async (token: string, transferId: string) => {
    return apiRequest(`/v1/transfers/${transferId}/cancel`, {
      method: 'POST',
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
  },

  getExchangeRates: async (token: string, from: string, to: string, amount: number) => {
    return apiRequest(`/v1/transfers/rates?from=${from}&to=${to}&amount=${amount}`, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
  },
};

interface Beneficiary {
  name: string;
  email: string;
  country: string;
  bankName: string;
  accountNumber: string;
  swiftCode: string;
}

// Beneficiary API
export const beneficiaryApi = {
  getBeneficiaries: async (token: string) => {
    return apiRequest('/v1/beneficiaries', {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
  },

  createBeneficiary: async (token: string, beneficiaryData: Beneficiary) => {
    return apiRequest('/v1/beneficiaries', {
      method: 'POST',
      headers: {
        Authorization: `Bearer ${token}`,
      },
      body: JSON.stringify(beneficiaryData),
    });
  },

  updateBeneficiary: async (token: string, beneficiaryId: string, beneficiaryData: Beneficiary) => {
    return apiRequest(`/v1/beneficiaries/${beneficiaryId}`, {
      method: 'PUT',
      headers: {
        Authorization: `Bearer ${token}`,
      },
      body: JSON.stringify(beneficiaryData),
    });
  },

  deleteBeneficiary: async (token: string, beneficiaryId: string) => {
    return apiRequest(`/v1/beneficiaries/${beneficiaryId}`, {
      method: 'DELETE',
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });
  },
};

// Health API
export const healthApi = {
  healthCheck: async () => {
    return apiRequest('/health');
  },

  readinessCheck: async () => {
    return apiRequest('/health/ready');
  },
};
