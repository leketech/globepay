# GlobePay Frontend

This is the frontend application for GlobePay, a fintech platform for international money transfers.

## Project Structure

```
src/
├── assets/              # Static assets (images, icons, etc.)
├── components/          # Reusable UI components
│   ├── auth/           # Authentication components (Login, Signup)
│   ├── common/         # Common components (ErrorBoundary, PrivateRoute)
│   ├── dashboard/      # Dashboard components
│   ├── layout/         # Layout components (Header, Layout)
│   ├── transactions/   # Transaction components
│   └── transfer/       # Transfer components
├── hooks/              # Custom React hooks
├── pages/              # Page components
├── services/           # API service layer
├── store/              # Redux store and slices
├── types/              # TypeScript type definitions
└── utils/              # Utility functions
```

## Components Completed

### Layout Components
- `Header.tsx` - Navigation header with user info and logout
- `Layout.tsx` - Main layout wrapper

### Dashboard Components
- `AccountSummary.tsx` - Displays user account balances

### Transaction Components
- `TransactionHistory.tsx` - Displays recent transactions

### Transfer Components
- `TransferForm.tsx` - Form for initiating money transfers

### Authentication Components
- `Login.tsx` - Login form
- `Signup.tsx` - User registration form

### Common Components
- `ErrorBoundary.tsx` - Error handling component
- `PrivateRoute.tsx` - Route protection component

## Utilities

### Formatting Utilities
- `format.ts` - Currency, date, and string formatting functions

### Validation Utilities
- `validation.ts` - Form validation functions

## Types

TypeScript interfaces for:
- User
- Account
- Transaction
- Transfer
- Redux state types

## Services

API service layer for:
- Authentication
- Transfer operations

## Store

Redux store with slices for:
- Authentication state
- Transfer state
- Transaction state

## Getting Started

1. Install dependencies:
   ```bash
   npm install
   ```

2. Start the development server:
   ```bash
   npm run dev
   ```

3. Build for production:
   ```bash
   npm run build
   ```

## Testing

Tests are implemented using Jest and React Testing Library. Test files are colocated with their corresponding components.

## Dependencies Needed

To run this application, you'll need to install the following dependencies:

```bash
npm install react react-dom react-router-dom redux react-redux @reduxjs/toolkit
npm install -D typescript @types/react @types/react-dom @types/node
npm install -D tailwindcss postcss autoprefixer
npm install -D @testing-library/react @testing-library/jest-dom @testing-library/user-event
npm install -D jest @types/jest
```

## Environment Variables

Create a `.env` file in the root of the frontend directory with the following variables:

```
VITE_API_URL=http://localhost:8080/api
VITE_APP_NAME=GlobePay
```