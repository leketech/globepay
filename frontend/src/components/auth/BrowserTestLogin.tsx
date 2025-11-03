import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { useAuth } from '../../hooks/useAuth';
import { authApi } from '../../services/api';

const BrowserTestLogin: React.FC = () => {
  const [email, setEmail] = useState('fresh_user@example.com');
  const [password, setPassword] = useState('password123');
  const [testResult, setTestResult] = useState<string>('');
  const [isTesting, setIsTesting] = useState(false);

  const navigate = useNavigate();
  const { isAuthenticated, user } = useAuth();

  const handleTestLogin = async () => {
    setIsTesting(true);
    setTestResult('');

    try {
      console.log('Testing login with:', { email, password });
      const result = await authApi.login(email, password);
      console.log('Login result:', result);

      // Store token and user in localStorage
      if (result.token) {
        localStorage.setItem('token', result.token);
        localStorage.setItem('user', JSON.stringify(result.user));
        setTestResult('SUCCESS: Login successful!');
      } else {
        setTestResult('ERROR: No token received');
      }

      // Navigate to dashboard after a short delay
      setTimeout(() => {
        navigate('/dashboard');
      }, 1000);
    } catch (err) {
      console.error('Login error:', err);
      if (err instanceof Error) {
        setTestResult(`ERROR: ${err.message || 'Login failed'}`);
      } else {
        setTestResult('ERROR: Login failed');
      }
    } finally {
      setIsTesting(false);
    }
  };

  const handleGoToDashboard = () => {
    navigate('/dashboard');
  };

  const handleGoToSignup = () => {
    navigate('/signup');
  };

  return (
    <div
      style={{
        padding: '20px',
        fontFamily: 'Arial, sans-serif',
        maxWidth: '500px',
        margin: '0 auto',
      }}
    >
      <h1>Browser Login Test</h1>
      <p>This page helps test the login functionality directly.</p>

      <div style={{ marginBottom: '15px' }}>
        <label>Email: </label>
        <input
          type="email"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          style={{ marginLeft: '10px', padding: '8px', width: '300px' }}
        />
      </div>

      <div style={{ marginBottom: '15px' }}>
        <label>Password: </label>
        <input
          type="password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          style={{ marginLeft: '10px', padding: '8px', width: '300px' }}
        />
      </div>

      <div style={{ marginBottom: '15px' }}>
        <button
          onClick={handleTestLogin}
          disabled={isTesting}
          style={{
            padding: '10px 20px',
            marginRight: '10px',
            backgroundColor: '#007bff',
            color: 'white',
            border: 'none',
            borderRadius: '4px',
            cursor: 'pointer',
          }}
        >
          {isTesting ? 'Testing...' : 'Test Login'}
        </button>

        <button
          onClick={handleGoToDashboard}
          style={{
            padding: '10px 20px',
            marginRight: '10px',
            backgroundColor: '#28a745',
            color: 'white',
            border: 'none',
            borderRadius: '4px',
            cursor: 'pointer',
          }}
        >
          Go to Dashboard
        </button>

        <button
          onClick={handleGoToSignup}
          style={{
            padding: '10px 20px',
            backgroundColor: '#ffc107',
            color: 'black',
            border: 'none',
            borderRadius: '4px',
            cursor: 'pointer',
          }}
        >
          Go to Signup
        </button>
      </div>

      {testResult && (
        <div
          style={{
            padding: '15px',
            marginTop: '20px',
            borderRadius: '4px',
            backgroundColor: testResult.startsWith('SUCCESS') ? '#d4edda' : '#f8d7da',
            color: testResult.startsWith('SUCCESS') ? '#155724' : '#721c24',
            border: `1px solid ${testResult.startsWith('SUCCESS') ? '#c3e6cb' : '#f5c6cb'}`,
          }}
        >
          <strong>{testResult}</strong>
        </div>
      )}

      {isAuthenticated && user && (
        <div
          style={{
            padding: '15px',
            marginTop: '20px',
            borderRadius: '4px',
            backgroundColor: '#d1ecf1',
            color: '#0c5460',
            border: '1px solid #bee5eb',
          }}
        >
          <strong>Currently Authenticated</strong>
          <p>
            User: {user.firstName} {user.lastName} ({user.email})
          </p>
        </div>
      )}
    </div>
  );
};

export default BrowserTestLogin;
