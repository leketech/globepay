import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { useDispatch, useSelector } from 'react-redux';
import { RootState, AppDispatch } from '../../store';
import { login, clearError } from '../../store/authSlice';
import { authApi } from '../../services/api';

const DebugLogin: React.FC = () => {
  const [email, setEmail] = useState('final_clean_test@example.com');
  const [password, setPassword] = useState('password123');
  const [apiResult, setApiResult] = useState<any>(null);
  const [apiError, setApiError] = useState<string | null>(null);
  const [loading, setLoading] = useState(false);
  
  const dispatch: AppDispatch = useDispatch();
  const navigate = useNavigate();
  const authState = useSelector((state: RootState) => state.auth);

  // Log auth state changes
  useEffect(() => {
    console.log('Auth state updated:', authState);
  }, [authState]);

  const handleDirectApiCall = async () => {
    setLoading(true);
    setApiError(null);
    setApiResult(null);
    
    try {
      console.log('Making direct API call with:', { email, password });
      const result = await authApi.login(email, password);
      console.log('Direct API call result:', result);
      setApiResult(result);
    } catch (err: any) {
      console.error('Direct API call error:', err);
      setApiError(err.message || 'API call failed');
    } finally {
      setLoading(false);
    }
  };

  const handleReduxLogin = async () => {
    setLoading(true);
    setApiError(null);
    setApiResult(null);
    
    try {
      console.log('Dispatching Redux login with:', { email, password });
      const resultAction = await dispatch(login({ email, password }));
      console.log('Redux login result action:', resultAction);
      
      if (login.fulfilled.match(resultAction)) {
        console.log('Redux login successful:', resultAction.payload);
        setApiResult(resultAction.payload);
      } else if (login.rejected.match(resultAction)) {
        console.log('Redux login rejected:', resultAction.payload);
        setApiError(resultAction.payload as string || 'Login failed');
      }
    } catch (err: any) {
      console.error('Redux login error:', err);
      setApiError(err.message || 'Login failed');
    } finally {
      setLoading(false);
    }
  };

  const handleClearError = () => {
    dispatch(clearError());
  };

  const handleNavigateToDashboard = () => {
    navigate('/dashboard');
  };

  const handleCheckLocalStorage = () => {
    console.log('LocalStorage token:', localStorage.getItem('token'));
    console.log('LocalStorage user:', localStorage.getItem('user'));
  };

  return (
    <div style={{ padding: '20px', fontFamily: 'Arial, sans-serif', maxWidth: '800px', margin: '0 auto' }}>
      <h1>Debug Login Component</h1>
      
      <div style={{ marginBottom: '20px', padding: '10px', backgroundColor: '#f0f0f0' }}>
        <h3>Current Auth State:</h3>
        <pre>{JSON.stringify(authState, null, 2)}</pre>
      </div>
      
      <div style={{ marginBottom: '20px' }}>
        <label>Email: </label>
        <input 
          type="email" 
          value={email} 
          onChange={(e) => setEmail(e.target.value)}
          style={{ marginLeft: '10px', padding: '5px', width: '300px' }}
        />
      </div>
      
      <div style={{ marginBottom: '20px' }}>
        <label>Password: </label>
        <input 
          type="password" 
          value={password} 
          onChange={(e) => setPassword(e.target.value)}
          style={{ marginLeft: '10px', padding: '5px', width: '300px' }}
        />
      </div>
      
      <div style={{ marginBottom: '20px' }}>
        <button 
          onClick={handleDirectApiCall} 
          disabled={loading}
          style={{ padding: '10px 20px', marginRight: '10px' }}
        >
          Direct API Call
        </button>
        
        <button 
          onClick={handleReduxLogin} 
          disabled={loading}
          style={{ padding: '10px 20px', marginRight: '10px' }}
        >
          Redux Login
        </button>
        
        <button 
          onClick={handleClearError}
          style={{ padding: '10px 20px', marginRight: '10px' }}
        >
          Clear Error
        </button>
        
        <button 
          onClick={handleCheckLocalStorage}
          style={{ padding: '10px 20px', marginRight: '10px' }}
        >
          Check LocalStorage
        </button>
        
        <button 
          onClick={handleNavigateToDashboard}
          style={{ padding: '10px 20px' }}
        >
          Go to Dashboard
        </button>
      </div>
      
      {apiError && (
        <div style={{ 
          marginTop: '20px', 
          padding: '10px', 
          backgroundColor: '#ffebee', 
          color: '#c62828',
          border: '1px solid #ffcdd2'
        }}>
          <h3>Error:</h3>
          <p>{apiError}</p>
        </div>
      )}
      
      {apiResult && (
        <div style={{ 
          marginTop: '20px', 
          padding: '10px', 
          backgroundColor: '#e8f5e9', 
          color: '#2e7d32',
          border: '1px solid #c8e6c9'
        }}>
          <h3>Result:</h3>
          <pre>{JSON.stringify(apiResult, null, 2)}</pre>
        </div>
      )}
    </div>
  );
};

export default DebugLogin;