import { AddIcon } from '@chakra-ui/icons';
import { Text, Flex, Button, Img, Box } from '@chakra-ui/react';
import { nanoid } from 'nanoid';
import { useEffect, useRef, useState } from 'react';
import { ReactSortable } from 'react-sortablejs';
import { uploadFiles } from '../apis/file';
import PhotoInputItem from './PhotoInputItem';

const ImageInput = ({ imageInfos, setImageInfos }) => {
  const inputRef = useRef();

  useEffect(() => {
    console.log(imageInfos);
  }, [imageInfos]);

  const uploadImages = async images => {
    const res = await uploadFiles(images.map(image => image.file));
    console.log(res);
    const uploadedImages = res.data.files;
    setImageInfos(prev =>
      prev.map(imageInfo => {
        const uploadedImage = uploadedImages.find(
          uploadedImage => uploadedImage.name === imageInfo.file.name
        );
        if (uploadedImage) {
          return {
            ...imageInfo,
            isUploading: false,
            id: uploadedImage.ID,
          };
        }
        return imageInfo;
      })
    );
  };

  const onImageUpload = e => {
    const targetFiles = Array.from(e.target.files).map(file => ({
      file: file,
      uuid: nanoid(),
      url: URL.createObjectURL(file),
      isUploading: true,
      id: 0,
    }));

    if (targetFiles.length + imageInfos.length > 5) {
      alert('최대 5장까지 업로드 가능합니다.');
      e.target.files = null;
      return;
    }

    setImageInfos(prev => [...prev, ...targetFiles]);

    uploadImages(targetFiles);

    e.target.files = null;
  };

  const onImageDelete = uuid => {
    setImageInfos(prev => prev.filter(imageInfo => imageInfo.uuid !== uuid));
  };

  return (
    <Flex flexDirection="column" gap="5px" width="100%">
      <Text fontWeight="bold">이미지 업로드</Text>
      <Text color="gray.600" fontSize="xs">
        최소 3장에서 최대 5장의 이미지를 업로드 해주세요.
      </Text>
      <input
        type="file"
        accept="image/*,*.heic"
        ref={inputRef}
        onChange={onImageUpload}
        multiple
        hidden
      />
      <Flex gap="8px" width={'100%'} align="center">
        <Button
          h="80px"
          w="80px"
          flexShrink={0}
          colorScheme="yellow"
          onClick={() => {
            inputRef.current.click();
          }}
          boxShadow="md"
        >
          <AddIcon />
        </Button>
        <ReactSortable
          list={imageInfos}
          setList={setImageInfos}
          animation={200}
          style={{
            display: 'flex',
            flexDirection: 'row',
            justifyContent: 'flex-start',
            padding: '8px 0px 8px',
            gap: '8px',
            width: '100%',
            overflowX: 'scroll',
          }}
        >
          {imageInfos.map((imageInfo, index) => {
            return (
              <PhotoInputItem
                index={index}
                key={imageInfo.uuid}
                src={imageInfo.url}
                isUploading={imageInfo.isUploading}
                onDelete={() => {
                  onImageDelete(imageInfo.uuid);
                }}
              />
            );
          })}
        </ReactSortable>
      </Flex>
    </Flex>
  );
};
export default ImageInput;
