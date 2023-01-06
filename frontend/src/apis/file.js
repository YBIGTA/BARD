import axiosInstance from './instance';

export const uploadFiles = async files => {
  const formData = new FormData();
  files.forEach(file => {
    console.log(file);
    formData.append('files[]', file);
  });

  try {
    const res = await axiosInstance.post(`/files/upload`, formData, {
      headers: {
        'Content-Type': 'multipart/form-data',
      },
    });
    return res;
  } catch (e) {
    console.log(e);
  }
};
