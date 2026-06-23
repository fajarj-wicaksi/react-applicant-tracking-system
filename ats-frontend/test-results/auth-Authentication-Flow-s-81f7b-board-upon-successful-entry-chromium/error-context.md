# Instructions

- Following Playwright test failed.
- Explain why, be concise, respect Playwright best practices.
- Provide a snippet of code with the fix, if possible.

# Test info

- Name: auth.spec.ts >> Authentication Flow >> should display the dashboard upon successful entry
- Location: e2e\auth.spec.ts:4:3

# Error details

```
Error: expect(locator).toBeVisible() failed

Locator:  locator('aside').getByText('TalentFlow')
Expected: visible
Received: hidden
Timeout:  5000ms

Call log:
  - Expect "toBeVisible" with timeout 5000ms
  - waiting for locator('aside').getByText('TalentFlow')
    13 × locator resolved to <span class="text-2xl font-bold bg-gradient-to-r from-primary to-blue-600 bg-clip-text text-transparent">TalentFlow</span>
       - unexpected value "hidden"

```

```yaml
- banner:
  - button
  - text: TalentFlow
  - button
  - text: Recruiter
- main:
  - heading "Dashboard" [level=1]
  - paragraph: Welcome to TalentFlow ATS
```

# Test source

```ts
  1  | import { test, expect } from '@playwright/test';
  2  | 
  3  | test.describe('Authentication Flow', () => {
  4  |   test('should display the dashboard upon successful entry', async ({ page }) => {
  5  |     await page.goto('/');
  6  | 
> 7  |     await expect(page.locator('aside').getByText('TalentFlow')).toBeVisible();
     |                                                                 ^ Error: expect(locator).toBeVisible() failed
  8  | 
  9  |     await expect(page.getByRole('heading', { name: 'Dashboard' })).toBeVisible();
  10 |     await expect(page.getByText('Welcome to TalentFlow ATS')).toBeVisible();
  11 |   });
  12 | 
  13 |   test('navigation links work correctly', async ({ page }) => {
  14 |     await page.goto('/');
  15 | 
  16 |     await page.getByRole('link', { name: 'Jobs' }).click();
  17 |     await expect(page.getByRole('heading', { name: 'Jobs' })).toBeVisible();
  18 | 
  19 |     await page.getByRole('link', { name: 'Candidates' }).click();
  20 |     await expect(page.getByRole('heading', { name: 'Candidates' })).toBeVisible();
  21 |   });
  22 | });
  23 | 
```