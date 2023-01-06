import axiosInstance from './instance';

export const googleSignIn = async ({ clientId, credential }) => {
  try {
    const res = await axiosInstance.post(
      `/auth/google`,
      {
        clientId,
        credential,
      },
      {
        withCredentials: true,
      }
    );

    return res;
  } catch (e) {
    console.log(e);
  }
};

export const logout = async () => {
  try {
    const res = await axiosInstance.get(`/auth/logout`);
    return res;
  } catch (e) {
    console.log(e);
  }
};

export const getSession = async () => {
  try {
    const res = await axiosInstance.get(`/auth/user`);

    return res;
  } catch (e) {
    console.log(e);
  }
};
