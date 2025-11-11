import { screen } from '@testing-library/react';
import { renderWithProviders } from '../../test-utils';
import { Header } from './Header';

// Mock the HeaderProfile component since it has complex dependencies
jest.mock('./HeaderProfile', () => ({
  HeaderProfile: () => <div data-testid="header-profile">Profile</div>,
}));

describe('Header Component', () => {
  it('should render without crashing', () => {
    renderWithProviders(<Header />);

    // Check that the header renders
    expect(screen.getByText('Globepay')).toBeInTheDocument();
    expect(screen.getByText('Dashboard')).toBeInTheDocument();
    expect(screen.getByText('Send Money')).toBeInTheDocument();
  });
});
