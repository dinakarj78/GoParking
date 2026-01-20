import type { ReactNode } from 'react'
import React from 'react'

export const metadata = {
  title: 'GoParking',
  description: 'GoParking app',
}

export default function RootLayout({ children }: { children: ReactNode }) {
  return (
    <html lang="en">
      <body>{children}</body>
    </html>
  )
}
