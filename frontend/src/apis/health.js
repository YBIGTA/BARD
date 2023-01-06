import axios from 'axios';
import axiosInstance from './instance';

const API_URL = process.env.REACT_APP_API_URL;

export const healthCheck = async () => {
  try {
    const res = await axiosInstance.get(`/health`);
    return res;
  } catch (e) {
    console.log(e);
  }
};
