import { test, expect } from '@playwright/test';

test.describe('Authentication Flow', () => {
  test('should display the dashboard upon successful entry', async ({ page }) => {
    await page.goto('/');

    await expect(page.getByText('TalentFlow').first()).toBeVisible();

    await expect(page.getByRole('heading', { name: 'Dashboard' })).toBeVisible();
    await expect(page.getByText('Welcome to TalentFlow ATS')).toBeVisible();
  });

  test('navigation links work correctly', async ({ page }) => {
    await page.goto('/');

    await page.getByRole('link', { name: 'Jobs' }).first().click();
    await expect(page.getByRole('heading', { name: 'Jobs' })).toBeVisible();

    await page.getByRole('link', { name: 'Candidates' }).first().click();
    await expect(page.getByRole('heading', { name: 'Candidates' })).toBeVisible();
  });
});
