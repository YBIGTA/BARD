import { Img } from '@chakra-ui/image';
import { Flex, Text } from '@chakra-ui/layout';
import moment from 'moment';

const StoryPreview = ({ data, onClick }) => {
  return (
    <Flex
      flexDirection="row"
      align="center"
      p="10px"
      gap="10px"
      isolation="isolate"
      w="calc(100vw - 48px)"
      maxW="calc(768px - 48px)"
      h="100px"
      border="1px solid"
      boxShadow="lg"
      borderRadius="5px"
      alignSelf="stretch"
      cursor="pointer"
      onClick={onClick}
    >
      <Flex
        flexDirection="column"
        align="flex-start"
        gap="10px"
        w="calc(100vw - 48px - 90px)"
        maxW="calc(768px - 48p - 90px)"
      >
        <Text fontWeight="bold">
          {data.title ? data.title : `Generated Story #${data.ID}`}
        </Text>
        <Text fontSize="sm" color="gray.500">
          {moment(data.createdAt).format('YYYY.MM.DD')}
        </Text>
      </Flex>
      <Img h="80px" w="80px" src={data.image_url} />
    </Flex>
  );
};

export default StoryPreview;
