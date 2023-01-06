import { Divider, Flex, Img, Text } from '@chakra-ui/react';
import { useNavigate } from 'react-router';
import BottomButton from '../components/BottomButton';
import bugSvg from '../assets/bug.svg';
import flowersDividerSvg from '../assets/flowersDivider.svg';
import tmpImg1 from '../assets/MoonStarsRocket.png';
import landingImg1 from '../assets/AboutFriendship.png';
import landingImg2 from '../assets/EverlastingMoment.png';
import { EmailIcon } from '@chakra-ui/icons';

function Landing() {
  const navigate = useNavigate();
  const onStart = () => {
    navigate('/login');
  };
  return (
    <>
      <Flex
        flexDirection="column"
        align="center"
        p="42px 24px 100px"
        gap="42px"
      >
        <Text textAlign="center" fontSize="6xl" fontFamily="Rochester">
          BARD
        </Text>
        <Text>
          BARD composes a new magical story based on your own creative
          characters with your selfies, landscapes, food, and any other daily
          photos.
        </Text>
        <Img w="68px" src={bugSvg} />
        <Flex flexDirection="column" align="center" gap="0px">
          <Text fontFamily="Aladin" fontSize="22px">
            Stories by BARD
          </Text>
          <Img
            w="calc(100vw - 48px)"
            maxW="calc(768px - 48px)"
            src={flowersDividerSvg}
          />
        </Flex>
        <Flex
          flexDirection="column"
          align="flex-start"
          p="12px 10px"
          gap="10px"
          border="0.7px solid"
          w="calc(100vw - 48px)"
          maxW="calc(768px - 48px)"
          borderRadius="5px"
        >
          <Text fontSize="2xl">About Friendship</Text>
          <Flex
            flexDirection="row"
            justify="center"
            align="center"
            p="0px 0px 4px"
            gap="20px"
          >
            <Img w="180px" src={landingImg1} />
            <Text fontSize="sm" textAlign="center">
              "예림과 영주는 친구들과 경기를 보며 즐거운 시간을 보냈답니다."
            </Text>
          </Flex>
        </Flex>
        <Flex
          flexDirection="column"
          align="flex-start"
          p="12px 10px"
          gap="10px"
          border="0.7px solid"
          w="calc(100vw - 48px)"
          maxW="calc(768px - 48px)"
          borderRadius="5px"
        >
          <Text fontSize="2xl">Everlasting Moment</Text>
          <Flex
            flexDirection="row"
            justify="center"
            align="center"
            p="0px 0px 4px"
            gap="20px"
          >
            <Img w="180px" src={landingImg2} />
            <Text fontSize="sm" textAlign="center">
              “수림은 고양이에게 먹을 것을 나눠 주고 둘도 없는 친구가
              되었습니다."
            </Text>
          </Flex>
        </Flex>
        <BottomButton onClick={onStart} title="Get Started" />
        <Flex flexDirection="column" align="center" gap="10px">
          <Divider w="calc(100vw - 48px)" maxW="calc(768px - 48px)" />
          <Text color="gray.500" fontSize="xs">
            2022 @ YBigta Team BARD
          </Text>
          <Text color="gray.500" fontSize="xs">
            Contact
          </Text>
          <Flex gap="10px">
            <EmailIcon color="gray.500" />
            <Text color="gray.500" fontSize="xs">
              tjguwns5757@gmail.com
            </Text>
          </Flex>
        </Flex>
      </Flex>
    </>
  );
}

export default Landing;
