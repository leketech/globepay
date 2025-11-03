import React from 'react';
import { Link } from 'react-router-dom';
import { HeaderProfile } from './HeaderProfile';

export const Header: React.FC = () => {
  return (
    <header className="flex items-center justify-between whitespace-nowrap border-b border-solid border-gray-200 dark:border-gray-700 px-10 py-3 bg-white dark:bg-[#182832]">
      <div className="flex items-center gap-4 text-primary">
        <div className="size-6">
          <svg fill="none" viewBox="0 0 48 48" xmlns="http://www.w3.org/2000/svg">
            <path d="M42.4379 44C42.4379 44 36.0744 33.9038 41.1692 24C46.8624 12.9336 42.2078 4 42.2078 4L7.01134 4C7.01134 4 11.6577 12.932 5.96912 23.9969C0.876273 33.9029 7.27094 44 7.27094 44L42.4379 44Z" fill="currentColor"></path>
          </svg>
        </div>
        <h2 className="text-lg font-bold leading-tight tracking-[-0.015em]">Globepay</h2>
      </div>
      <nav className="flex flex-1 justify-center items-center gap-6">
        <Link to="/dashboard" className="text-sm font-medium leading-normal text-gray-600 dark:text-gray-300 hover:text-primary dark:hover:text-primary transition-colors flex items-center gap-1">
          <span className="material-symbols-outlined text-base">dashboard</span>
          <span>Dashboard</span>
        </Link>
        <Link to="/transfer" className="text-sm font-medium leading-normal text-gray-600 dark:text-gray-300 hover:text-primary dark:hover:text-primary transition-colors flex items-center gap-1">
          <span className="material-symbols-outlined text-base">send</span>
          <span>Send Money</span>
        </Link>
        <Link to="/recipients" className="text-sm font-medium leading-normal text-gray-600 dark:text-gray-300 hover:text-primary dark:hover:text-primary transition-colors flex items-center gap-1">
          <span className="material-symbols-outlined text-base">group</span>
          <span>Recipients</span>
        </Link>
        <Link to="/history" className="text-sm font-medium leading-normal text-gray-600 dark:text-gray-300 hover:text-primary dark:hover:text-primary transition-colors flex items-center gap-1">
          <span className="material-symbols-outlined text-base">history</span>
          <span>History</span>
        </Link>
        <Link to="/settings" className="text-sm font-medium leading-normal text-gray-600 dark:text-gray-300 hover:text-primary dark:hover:text-primary transition-colors flex items-center gap-1">
          <span className="material-symbols-outlined text-base">settings</span>
          <span>Settings</span>
        </Link>
      </nav>
      <div className="flex items-center gap-2">
        <HeaderProfile />
      </div>
    </header>
  );
};