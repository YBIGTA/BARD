import React from 'react';
import { ChakraProvider, Box, extendTheme } from '@chakra-ui/react';
import { Route, Routes } from 'react-router-dom';

import Landing from './pages/Landing';
import Login from './pages/Login';
import Home from './pages/Home';
import SignUpPolicy from './pages/SignUpPolicy';
import StoryView from './pages/StoryView';
import StoryForm from './pages/StoryForm';
import SignUpName from './pages/SignUpName';
import paperBackground from '../src/assets/SmoothNotebookPaper_exposed3.png';
import { GoogleOAuthProvider } from '@react-oauth/google';

const customTheme = extendTheme({
  colors: {
    yellow: {
      400: '#F7DE84', // Override the default yellow color (500) with red
    },
  },
});

function App() {
  return (
    <GoogleOAuthProvider clientId={process.env.REACT_APP_GOOGLE_CLIENT_ID}>
      <ChakraProvider theme={customTheme}>
        <Box boxSizing="border-box">
          <Box
            m={'0 auto'} // 어플리케이션 영역을 중앙에 위치시킴
            minH="100vh"
            w="100vw" // 어플리케이션 영역 너비를 뷰포트 너비로 설정
            maxW="48rem" // 어플리케이션 영역 너비가 48rem(768px)을 넘지 않도록 설정
            overflow="hidden"
            boxShadow="md"
            rounded="md"
            backgroundImage={paperBackground}
            // filter="brightness(108%)"
          >
            <Routes>
              <Route path="/" element={<Landing />} />
              <Route path="/login" element={<Login />} />
              <Route path="/signup/policy" element={<SignUpPolicy />} />
              <Route path="/signup/user" element={<SignUpName />} />
              <Route path="/home" element={<Home />} />
              <Route path="/story/new" element={<StoryForm />} />
              <Route path="/story/:storyId" element={<StoryView />} />
            </Routes>
          </Box>
        </Box>
      </ChakraProvider>
    </GoogleOAuthProvider>
  );
}

export default App;
