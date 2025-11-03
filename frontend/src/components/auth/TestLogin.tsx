import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { useDispatch } from 'react-redux';
import { login } from '../../store/authSlice';

const TestLogin: React.FC = () => {
  const [email, setEmail] = useState('final_clean_test@example.com');
  const [password, setPassword] = useState('password123');
  const [result, setResult] = useState<any>(null);
  const [error, setError] = useState<string | null>(null);
  const [loading, setLoading] = useState(false);
  
  const dispatch = useDispatch();
  const navigate = useNavigate();

  const handleTestLogin = async () => {
    setLoading(true);
    setError(null);
    setResult(null);
    
    try {
      console.log('Dispatching login action with:', { email, password });
      const actionResult: any = await dispatch(login({ email, password }) as any);
      console.log('Login action result:', actionResult);
      
      if (login.fulfilled.match(actionResult)) {
        setResult(actionResult.payload);
        console.log('Login successful!');
      } else if (login.rejected.match(actionResult)) {
        setError(actionResult.error?.message || 'Login failed');
        console.log('Login rejected:', actionResult.error);
      }
    } catch (err) {
      console.error('Login error:', err);
      setError(err instanceof Error ? err.message : 'Login failed');
    } finally {
      setLoading(false);
    }
  };

  const handleNavigateToDashboard = () => {
    navigate('/dashboard');
  };

  return (
    <div style={{ padding: '20px', fontFamily: 'Arial, sans-serif' }}>
      <h1>Test Login Component</h1>
      
      <div style={{ marginBottom: '20px' }}>
        <label>Email: </label>
        <input 
          type="email" 
          value={email} 
          onChange={(e) => setEmail(e.target.value)}
          style={{ marginLeft: '10px', padding: '5px' }}
        />
      </div>
      
      <div style={{ marginBottom: '20px' }}>
        <label>Password: </label>
        <input 
          type="password" 
          value={password} 
          onChange={(e) => setPassword(e.target.value)}
          style={{ marginLeft: '10px', padding: '5px' }}
        />
      </div>
      
      <button 
        onClick={handleTestLogin} 
        disabled={loading}
        style={{ padding: '10px 20px', marginRight: '10px' }}
      >
        {loading ? 'Logging in...' : 'Test Login'}
      </button>
      
      <button 
        onClick={handleNavigateToDashboard}
        style={{ padding: '10px 20px' }}
      >
        Go to Dashboard
      </button>
      
      {error && (
        <div style={{ 
          marginTop: '20px', 
          padding: '10px', 
          backgroundColor: '#ffebee', 
          color: '#c62828',
          border: '1px solid #ffcdd2'
        }}>
          <h3>Error:</h3>
          <p>{error}</p>
        </div>
      )}
      
      {result && (
        <div style={{ 
          marginTop: '20px', 
          padding: '10px', 
          backgroundColor: '#e8f5e9', 
          color: '#2e7d32',
          border: '1px solid #c8e6c9'
        }}>
          <h3>Login Result:</h3>
          <pre>{JSON.stringify(result, null, 2)}</pre>
        </div>
      )}
    </div>
  );
};

export default TestLogin;