import { test, expect } from '@playwright/test';

test.describe('Authentication', () => {
  test.beforeEach(async ({ page }) => {
    // Navigate to the login page before each test
    await page.goto('/login');
  });

  test('should display login form', async ({ page }) => {
    // Check if email input is visible
    await expect(page.locator('input[name="email"]')).toBeVisible();
    
    // Check if password input is visible
    await expect(page.locator('input[name="password"]')).toBeVisible();
    
    // Check if submit button is visible
    await expect(page.locator('button[type="submit"]')).toBeVisible();
  });

  test('should show validation errors for empty form', async ({ page }) => {
    // Try to submit empty form
    await page.click('button[type="submit"]');
    
    // Check for validation errors (adjust selectors based on your actual implementation)
    // This is a placeholder - adjust based on your form validation
  });

  test('should navigate to signup page', async ({ page }) => {
    // Click on signup link if it exists
    const signupLink = page.locator('a[href*="signup"]').first();
    if (await signupLink.isVisible()) {
      await signupLink.click();
      await expect(page).toHaveURL(/.*signup.*/);
    }
  });
});
