// Simple API service for Globepay
import { Transfer } from '../types';

const API_BASE_URL = '/api'; // Use relative path to leverage Vite proxy

// Helper function for API requests
const apiRequest = async (endpoint: string, options: RequestInit = {}) => {
  // For health checks, we need to use the full URL
  const isHealthCheck = endpoint.startsWith('/health');
  const url = isHealthCheck ? `http://localhost:8080${endpoint}` : `${API_BASE_URL}${endpoint}`;

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

    if (!response.ok) {
      const errorData = await response.json().catch(() => ({}));
      console.error('API request failed:', {
        status: response.status,
        statusText: response.statusText,
        errorData,
      });
      throw new Error(
        `HTTP error! status: ${response.status}, message: ${JSON.stringify(errorData)}`
      );
    }

    const data = await response.json();
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
    return apiRequest('/v1/auth/login', {
      method: 'POST',
      body: JSON.stringify({ email, password }),
    });
  },

  register: async (userData: RegisterUserData) => {
    // Filter out empty optional fields
    const filteredData = filterEmptyFields(userData);
    return apiRequest('/v1/auth/register', {
      method: 'POST',
      body: JSON.stringify(filteredData),
    });
  },

  refreshToken: async (refreshToken: string) => {
    return apiRequest('/v1/auth/refresh', {
      method: 'POST',
      body: JSON.stringify({ refreshToken }),
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
