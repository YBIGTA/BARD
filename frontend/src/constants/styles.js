import { extendTheme } from '@chakra-ui/react';

const theme = extendTheme({
  fonts: {
    heading: `'pretendard', sans-serif`,
    body: `'pretendard', sans-serif`,
  },
  colors: {
    yellow: {
      100: 'rgba(236, 201, 75, 0.7)',
    },
  },
});

export default theme;
