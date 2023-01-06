import {
  Box,
  Button,
  Flex,
  Modal,
  ModalBody,
  ModalContent,
  ModalFooter,
  ModalOverlay,
  Text,
} from '@chakra-ui/react';
import { useState } from 'react';
import Lottie from 'react-lottie';
import BottomButton from '../components/BottomButton';
import CharacterInput from '../components/CharacterInput';
import Header from '../components/Header';
import ImageInput from '../components/ImageInput';
import MakingStoryLottieData from '../assets/MakingStoryLottie.json';
import MakingCompleteLottieData from '../assets/MakingCompleteLottie.json';
import { ArrowRightIcon } from '@chakra-ui/icons';
import { useNavigate } from 'react-router';
import { createStory } from '../apis/story';

function StoryForm() {
  const [characters, setCharacters] = useState([]);
  const [imageInfos, setImageInfos] = useState([]);

  const [isLoading, setIsLoading] = useState(false);
  const [createdStoryId, setCreatedStoryId] = useState(null);

  const onSubmit = async () => {
    const imageIds = imageInfos.map(imageInfo => imageInfo.id);

    setIsLoading(true);
    const res = await createStory({
      characters: characters,
      image_ids: imageIds,
    });

    console.log(res);

    if (res.status === 201) {
      setCreatedStoryId(res.data.story_id);
    } else {
      alert('스토리 생성에 실패했습니다.');
      setIsLoading(false);
    }
  };

  return (
    <>
      <Flex
        w="100%"
        flexDirection="column"
        p="0px"
        flexGrow="1"
        flexShrink="1"
        overflow="auto"
      >
        <Header title="스토리 만들기" isBack />
        <Flex
          flexDirection="column"
          justify="space-between"
          alignItems="center"
          p="60px 24px 40px"
        >
          <Flex
            flexDirection="column"
            align="flex-start"
            padding="0px"
            gap="32px"
            w="calc(100vw - 48px)"
            maxW="calc(768px - 48px)"
          >
            <ImageInput imageInfos={imageInfos} setImageInfos={setImageInfos} />
            <CharacterInput
              characters={characters}
              setCharacters={setCharacters}
            />
          </Flex>
          <BottomButton title="입력 완료" onClick={onSubmit} />
        </Flex>
      </Flex>
      <LoadingModal isOpen={isLoading} storyId={createdStoryId} />
    </>
  );
}

const LoadingModal = ({ isOpen, storyId }) => {
  const navigate = useNavigate();
  const isStoryMakingDone = storyId !== null;

  return (
    <Modal isCentered closeOnOverlayClick={false} isOpen={isOpen}>
      <ModalOverlay />
      <ModalContent>
        <ModalBody transition="all ease-in-out">
          <Box w="200" h="200">
            <Lottie
              options={{
                loop: true,
                autoplay: true,
                animationData: isStoryMakingDone
                  ? MakingCompleteLottieData
                  : MakingStoryLottieData,
              }}
            />
          </Box>
          <ModalFooter>
            {isStoryMakingDone ? (
              <Flex direction="column" margin="0 auto" gap="30px">
                <Text fontWeight="extrabold" fontSize="2xl" margin="0 auto">
                  스토리 생성 완료
                </Text>
                <Button
                  colorScheme="yellow"
                  leftIcon={<ArrowRightIcon />}
                  margin="0 auto"
                  fontSize="sm"
                  size="md"
                  onClick={() => {
                    navigate(`/story/${storyId}`);
                  }}
                >
                  확인해 보세요!
                </Button>
              </Flex>
            ) : (
              <Text fontWeight="bold" fontSize="xl" margin="0 auto">
                스토리 생성중...
              </Text>
            )}
          </ModalFooter>
        </ModalBody>
      </ModalContent>
    </Modal>
  );
};

export default StoryForm;
