# Instructions

- Following Playwright test failed.
- Explain why, be concise, respect Playwright best practices.
- Provide a snippet of code with the fix, if possible.

# Test info

- Name: auth.spec.ts >> Authentication Flow >> navigation links work correctly
- Location: e2e\auth.spec.ts:13:3

# Error details

```
Test timeout of 30000ms exceeded.
```

```
Error: locator.click: Test timeout of 30000ms exceeded.
Call log:
  - waiting for getByRole('link', { name: 'Jobs' })

```

# Page snapshot

```yaml
- generic [ref=e4]:
  - banner [ref=e5]:
    - generic [ref=e6]:
      - button [ref=e7]:
        - img [ref=e8]
      - generic [ref=e9]: TalentFlow
    - generic [ref=e10]:
      - button [ref=e11]:
        - img [ref=e12]
      - generic [ref=e15]:
        - img [ref=e17]
        - generic [ref=e20]: Recruiter
  - main [ref=e21]:
    - generic [ref=e22]:
      - heading "Dashboard" [level=1] [ref=e23]
      - paragraph [ref=e24]: Welcome to TalentFlow ATS
```

# Test source

```ts
  1  | import { test, expect } from '@playwright/test';
  2  | 
  3  | test.describe('Authentication Flow', () => {
  4  |   test('should display the dashboard upon successful entry', async ({ page }) => {
  5  |     await page.goto('/');
  6  | 
  7  |     await expect(page.locator('aside').getByText('TalentFlow')).toBeVisible();
  8  | 
  9  |     await expect(page.getByRole('heading', { name: 'Dashboard' })).toBeVisible();
  10 |     await expect(page.getByText('Welcome to TalentFlow ATS')).toBeVisible();
  11 |   });
  12 | 
  13 |   test('navigation links work correctly', async ({ page }) => {
  14 |     await page.goto('/');
  15 | 
> 16 |     await page.getByRole('link', { name: 'Jobs' }).click();
     |                                                    ^ Error: locator.click: Test timeout of 30000ms exceeded.
  17 |     await expect(page.getByRole('heading', { name: 'Jobs' })).toBeVisible();
  18 | 
  19 |     await page.getByRole('link', { name: 'Candidates' }).click();
  20 |     await expect(page.getByRole('heading', { name: 'Candidates' })).toBeVisible();
  21 |   });
  22 | });
  23 | 
```