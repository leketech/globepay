import { screen } from '@testing-library/react';
import { renderWithProviders } from '../../test-utils';
import Login from './Login';

// Mock the authService to avoid API calls
jest.mock('../../services/auth.service', () => ({
  authService: {
    login: jest.fn(),
    getCurrentUser: jest.fn(),
    getToken: jest.fn(),
    isAuthenticated: jest.fn(),
  }
}));

describe('Login Component', () => {
  it('should render without crashing', () => {
    renderWithProviders(<Login />);
    
    // Check that the login form renders
    expect(screen.getByText('Welcome Back')).toBeInTheDocument();
    expect(screen.getByLabelText('Email Address')).toBeInTheDocument();
    // Use a different approach to select the password input field
    expect(screen.getByPlaceholderText('Enter your password')).toBeInTheDocument();
  });
});