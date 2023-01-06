import { Button, Flex, Text } from '@chakra-ui/react';
import { useNavigate } from 'react-router';
import Header from '../components/Header';
import { useEffect, useState } from 'react';
import BottomButton from '../components/BottomButton';
import StoryPreview from '../components/StoryPreview';
import { getStories } from '../apis/story';
import { healthCheck } from '../apis/health';

function Home() {
  const navigate = useNavigate();
  const [stories, setStories] = useState([]);

  useEffect(() => {
    (async () => {
      const res = await getStories();
      console.log(res);
      if (res.status === 200) {
        setStories(res.data.stories);
      }
    })();
  }, []);

  return (
    <>
      <Flex
        // pt="3.5rem"
        w="100%"
        flexDirection="column"
        flexGrow="1"
        flexShrink="1"
        flexBasis="0%"
        overflow="auto"
      >
        <Header title="BARD" />
        <Flex
          flexDirection="column"
          alignItems="flex-start"
          p="40px 24px 10px"
          gap="24px"
        >
          <Flex justify="center" align="flex-start">
            <Text fontWeight="700" fontSize="4xl" color="yellow.500">
              Your Stories
            </Text>
          </Flex>
          <Flex
            flexDirection="column"
            align="flex-start"
            gap="16px"
            minHeight="100vh"
          >
            <Text fontWeight="semibold">Recent Stories</Text>
            {stories.length === 0 && (
              <Text fontWeight="semibold" color="gray.500">
                아직 스토리가 없습니다.
              </Text>
            )}
            {stories.map(item => (
              <StoryPreview
                key={item.ID}
                data={item}
                onClick={() => navigate(`/story/${item.ID}`)}
              />
            ))}
          </Flex>
          <BottomButton
            title="스토리 만들기"
            onClick={() => {
              navigate('/story/new');
            }}
          />
        </Flex>
      </Flex>
    </>
  );
}

export default Home;
