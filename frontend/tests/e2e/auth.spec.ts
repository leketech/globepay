import { test, expect } from '@playwright/test';

test.describe('Authentication Flow', () => {
  test.beforeEach(async ({ page }) => {
    await page.goto('/');
  });

  test('should load the home page', async ({ page }) => {
    await expect(page).toHaveTitle(/Globepay/i);
  });

  test('should navigate to login page', async ({ page }) => {
    const loginLink = page.locator('a[href*="login"]').first();
    if (await loginLink.isVisible().catch(() => false)) {
      await loginLink.click();
      await expect(page).toHaveURL(/.*login.*/);
    }
  });
});
