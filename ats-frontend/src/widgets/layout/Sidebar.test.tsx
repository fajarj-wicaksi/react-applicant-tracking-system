import { render, screen } from '@testing-library/react'
import { Sidebar } from './Sidebar'
import { describe, it, expect } from 'vitest'
import { BrowserRouter } from 'react-router-dom'

describe('Sidebar component', () => {
  const renderSidebar = () => {
    render(
      <BrowserRouter>
        <Sidebar />
      </BrowserRouter>
    )
  }

  it('renders the application title', () => {
    renderSidebar()
    expect(screen.getByText('TalentFlow')).toBeInTheDocument()
  })

  it('renders all navigation links', () => {
    renderSidebar()
    expect(screen.getByText('Dashboard')).toBeInTheDocument()
    expect(screen.getByText('Jobs')).toBeInTheDocument()
    expect(screen.getByText('Candidates')).toBeInTheDocument()
    expect(screen.getByText('Interviews')).toBeInTheDocument()
    expect(screen.getByText('Tasks')).toBeInTheDocument()
  })
})
