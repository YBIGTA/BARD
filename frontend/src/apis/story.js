import axiosInstance from './instance';

export const getStories = async () => {
  try {
    const res = await axiosInstance.get(`/stories`, {
      withCredentials: true,
    });
    return res;
  } catch (e) {
    console.log(e);
  }
};

export const getStoryById = async story_id => {
  try {
    const res = await axiosInstance.get(`/stories/${story_id}`);
    return res;
  } catch (e) {
    console.log(e);
  }
};

export const updateStoryTitle = async (story_id, title) => {
  try {
    const res = await axiosInstance.patch(`/stories/${story_id}/title`, {
      title: title,
    });
    return res;
  } catch (e) {
    console.log(e);
  }
};

export const createStory = async payload => {
  try {
    const res = await axiosInstance.post(`/stories`, {
      characters: payload.characters,
      image_ids: payload.image_ids,
    });
    return res;
  } catch (e) {
    console.log(e);
  }
};
