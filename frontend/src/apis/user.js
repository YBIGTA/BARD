import axiosInstance from './instance';

export const signupUser = async payload => {
  try {
    const res = await axiosInstance.post(`/users`, {
      email: payload.email,
      name: payload.name,
      social_id: payload.social_id,
    });
    return res;
  } catch (e) {
    console.log(e);
  }
};
