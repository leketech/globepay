import { screen } from '@testing-library/react';
import { renderWithProviders } from '../../test-utils';

describe('Simple Test', () => {
  it('should pass', () => {
    renderWithProviders(<div>Hello World</div>);
    expect(screen.getByText('Hello World')).toBeInTheDocument();
  });
});
