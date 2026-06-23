import { render, screen } from '@testing-library/react'
import { Header } from './Header'
import { describe, it, expect } from 'vitest'

describe('Header component', () => {
  it('renders the application title', () => {
    render(<Header />)
    expect(screen.getByText('TalentFlow')).toBeInTheDocument()
  })

  it('renders the user profile section', () => {
    render(<Header />)
    expect(screen.getByText('Recruiter')).toBeInTheDocument()
  })
})
