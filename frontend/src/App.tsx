import React, { useEffect } from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import { Provider } from 'react-redux';
import { store } from './store';
import { Layout } from './components/layout/Layout';
import { PrivateRoute } from './components/common/PrivateRoute';
import Welcome from './pages/Welcome';
import Dashboard from './pages/Dashboard';
import Transfer from './pages/Transfer';
import History from './pages/History';
import Recipients from './pages/Recipients';
import Profile from './pages/Profile';
import Settings from './pages/Settings';
import Login from './components/auth/Login';
import Signup from './components/auth/Signup';
import TestLogin from './components/auth/TestLogin';
import DebugLogin from './components/auth/DebugLogin';
import BrowserTestLogin from './components/auth/BrowserTestLogin';

const App: React.FC = () => {
  useEffect(() => {
    // Apply saved theme or system preference on app load
    const savedTheme = localStorage.getItem('theme') as 'light' | 'dark' | 'system' | null;
    
    const applyTheme = (theme: 'light' | 'dark' | 'system') => {
      if (theme === 'system') {
        const systemPrefersDark = window.matchMedia('(prefers-color-scheme: dark)').matches;
        if (systemPrefersDark) {
          document.documentElement.classList.add('dark');
        } else {
          document.documentElement.classList.remove('dark');
        }
      } else if (theme === 'dark') {
        document.documentElement.classList.add('dark');
      } else {
        document.documentElement.classList.remove('dark');
      }
    };
    
    if (savedTheme) {
      applyTheme(savedTheme);
    } else {
      // Default to system preference
      applyTheme('system');
    }
    
    // Listen for system theme changes
    const mediaQuery = window.matchMedia('(prefers-color-scheme: dark)');
    const handleSystemThemeChange = (e: MediaQueryListEvent) => {
      const currentTheme = localStorage.getItem('theme') as 'light' | 'dark' | 'system' | null;
      // Only apply system theme changes if user has selected 'system'
      if (!currentTheme || currentTheme === 'system') {
        if (e.matches) {
          document.documentElement.classList.add('dark');
        } else {
          document.documentElement.classList.remove('dark');
        }
      }
    };
    
    mediaQuery.addEventListener('change', handleSystemThemeChange);
    
    return () => {
      mediaQuery.removeEventListener('change', handleSystemThemeChange);
    };
  }, []);

  return (
    <Provider store={store}>
      <Router>
        <Routes>
          <Route path="/" element={<Welcome />} />
          <Route path="/login" element={<Login />} />
          <Route path="/test-login" element={<TestLogin />} />
          <Route path="/debug-login" element={<DebugLogin />} />
          <Route path="/browser-test-login" element={<BrowserTestLogin />} />
          <Route path="/signup" element={<Signup />} />
          <Route path="/dashboard" element={
            <PrivateRoute>
              <Layout>
                <Dashboard />
              </Layout>
            </PrivateRoute>
          } />
          <Route path="/transfer" element={
            <PrivateRoute>
              <Layout>
                <Transfer />
              </Layout>
            </PrivateRoute>
          } />
          <Route path="/history" element={
            <PrivateRoute>
              <Layout>
                <History />
              </Layout>
            </PrivateRoute>
          } />
          <Route path="/recipients" element={
            <PrivateRoute>
              <Layout>
                <Recipients />
              </Layout>
            </PrivateRoute>
          } />
          <Route path="/profile" element={
            <PrivateRoute>
              <Layout>
                <Profile />
              </Layout>
            </PrivateRoute>
          } />
          <Route path="/settings" element={
            <PrivateRoute>
              <Layout>
                <Settings />
              </Layout>
            </PrivateRoute>
          } />
        </Routes>
      </Router>
    </Provider>
  );
};

export default App;