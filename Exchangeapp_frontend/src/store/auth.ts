import { defineStore } from 'pinia';
import { ref, computed } from 'vue';
import axios from '../axios';

export const useAuthStore = defineStore('auth', () => {
  const token = ref<string | null>(localStorage.getItem('token'));

  const isAuthenticated = computed(() => !!token.value);

  const login = async (username: string, password: string) => {
    try {
      const response = await axios.post('/auth/login', { username, password });
      token.value = response.data.token;
      localStorage.setItem('token', token.value || '');
    } catch (error) {
      throw new Error(`Login failed! ${error}`);
    }
  };

  const register = async (username: string, password: string) => {
    try {
      const response = await axios.post('/auth/register', { username, password });
      token.value = response.data.token;
      localStorage.setItem('token', token.value || '');
    } catch (error) {
      throw new Error(`Register failed! ${error}`);
    }
  };

  const logout = () => {
    token.value = null;
    localStorage.removeItem('token');
  };

  return {
    token,
    isAuthenticated,
    login,
    register,
    logout
  };
});
