import { ChevronLeftIcon } from '@chakra-ui/icons';
import { Flex, Text } from '@chakra-ui/react';
import { useNavigate } from 'react-router';

function Header({ title, isBack }) {
  let navigate = useNavigate();
  const onGoBack = () => {
    navigate(-1);
  };

  return (
    <Flex
      boxSizing="border-box"
      flexDirection="row"
      alignItems="center"
      justify="center"
      p="16px 10px"
      gap="8px"
      borderBottom="2px solid #000000"
      isolation="isolate"
    >
      {isBack && (
        <ChevronLeftIcon
          position="absolute"
          left="16px"
          h="24px"
          w="24px"
          onClick={onGoBack}
        />
      )}
      <Text fontSize="md" fontWeight="bold">
        {title}
      </Text>
    </Flex>
  );
}

export default Header;
