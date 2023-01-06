import {
  Button,
  Flex,
  FormControl,
  FormLabel,
  Image,
  Img,
  Input,
  Text,
} from '@chakra-ui/react';
import { useEffect, useState } from 'react';
import Header from '../components/Header';
import MoonStarsRocketImg from '../assets/MoonStarsRocket.png';
import { ArrowRightIcon } from '@chakra-ui/icons';
import flowersDividerSvg from '../assets/flowersDivider.svg';
import useAppearSentencesOnScroll from '../hooks/useAppearSentencesOnScroll';
import { getStoryById, updateStoryTitle } from '../apis/story';
import { Navigate, useNavigate, useParams } from 'react-router';

function StoryView() {
  const [sentences, setSentences] = useState([]);
  const [image, setImage] = useState('');

  const { sentenceRefs } = useAppearSentencesOnScroll();

  const { storyId } = useParams();
  const navigate = useNavigate();

  useEffect(() => {
    (async () => {
      const res = await getStoryById(storyId);
      if (res.status === 200) {
        const story = res.data.story;
        setSentences(
          story.body
            .split(/'(".*?")\s|\.\s|\?\s/) //이게 ios에서 되면서 + 온점은 걍 아예 없이 출력.
            .filter(Boolean)
        );
        setImage(story.image_url);
      }
    })();
  }, [image, storyId]);

  const [title, setTitle] = useState('');
  const onTitleSubmit = async () => {
    if (title.trim() === '') {
      return;
    }
    setTitle('');
    const res = await updateStoryTitle(storyId, title);
    if (res.status === 201) {
      alert('제목이 저장되었습니다.');
      navigate('/home');
    } else {
      alert('제목 변경에 실패했습니다.');
    }
  };

  return (
    <>
      <Flex w="100%" flexDirection="column">
        <Header title="Read Story" isBack />
        <Flex
          flexDirection="column"
          alignItems="center"
          p="40px 24px 10px"
          margin="0 auto"
          padding="2rem"
          gap="42px"
          overflowWrap="break-word"
        >
          <Text textAlign="center" fontSize="6xl" fontFamily="Rochester">
            BARD
          </Text>
          <Img w="90%" src={MoonStarsRocketImg} />
          <Text fontFamily="Aladin" fontSize="22px">
            Scroll down slowly....
          </Text>
          {sentences.map((item, index) => (
            <Text
              fontSize="xl"
              ref={el => el && sentenceRefs.current.push(el)}
              transition="all 0.25s cubic-bezier(0.4, 0, 0.2, 1)"
              opacity="0"
              filter="blur(5px)"
              transitionDelay="all 3s"
              marginBottom="3.2rem"
              key={index}
              textAlign="center"
              wordBreak="keep-all"
            >
              {item}
            </Text>
          ))}
          <Image src={flowersDividerSvg} />
          <Image src={image} />
          <Flex flexDirection="column" align="center">
            <Text fontFamily="Aladin" fontSize="22px">
              Stories by BARD
            </Text>
            <Image transform="scaleY(-1)" src={flowersDividerSvg} />
          </Flex>

          <Flex
            flexDirection="column"
            align="flex-start"
            p="12px"
            w="100%"
            minH="80px"
            bgColor="#FCF6E2"
            borderRadius="5px"
          >
            <FormControl isRequired>
              <FormLabel>Enter your title:</FormLabel>
              <Input
                value={title}
                onChange={e => setTitle(e.target.value)}
                borderColor="yellow.800"
                variant="flushed"
                placeholder="제목을 입력하세요"
              />
            </FormControl>
          </Flex>
          <Button
            colorScheme="yellow"
            w="100%"
            gap="8px"
            border="0.7px solid"
            borderRadius="5px"
            onClick={onTitleSubmit}
          >
            <Text>저장하기</Text>
            <ArrowRightIcon boxSize={3} />
          </Button>
        </Flex>
      </Flex>
    </>
  );
}

export default StoryView;
