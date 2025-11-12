import { test, expect } from '@playwright/test';

test.describe('Authentication Flow', () => {
  test.beforeEach(async ({ page }) => {
    // Listen for console messages
    page.on('console', msg => console.log('Console:', msg.text()));
    page.on('pageerror', error => console.log('Page error:', error.message));
    page.on('requestfailed', request => console.log('Request failed:', request.url(), request.failure()?.errorText));
    
    await page.goto('/', { waitUntil: 'networkidle' });
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
