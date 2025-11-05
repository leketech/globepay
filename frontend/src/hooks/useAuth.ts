import { useSelector, useDispatch } from 'react-redux';
import { useCallback } from 'react';
import { RootState, AppDispatch } from '../store';
import { login, signup, logout } from '../store/authSlice';
import { LoginRequest, SignupRequest } from '../services/auth.service';

export const useAuth = () => {
  const dispatch = useDispatch<AppDispatch>();
  const { user, isAuthenticated, loading, error } = useSelector((state: RootState) => state.auth);

  const handleLogin = useCallback(
    (credentials: LoginRequest) => {
      console.log('useAuth login called with credentials:', credentials);
      return dispatch(login(credentials));
    },
    [dispatch]
  );

  const handleSignup = useCallback(
    (userData: SignupRequest) => {
      console.log('useAuth signup called with userData:', userData);
      return dispatch(signup(userData));
    },
    [dispatch]
  );

  const handleLogout = useCallback(() => {
    return dispatch(logout());
  }, [dispatch]);

  return {
    user,
    isAuthenticated,
    loading,
    error,
    login: handleLogin,
    signup: handleSignup,
    logout: handleLogout,
  };
};
