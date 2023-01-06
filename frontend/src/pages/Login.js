import { Flex, Image, Text } from '@chakra-ui/react';
import { useNavigate } from 'react-router';
import Header from '../components/Header';
import ViolinIcon from '../assets/violin.png';
import { GoogleLogin } from '@react-oauth/google';
import { googleSignIn } from '../apis/auth';

function Login() {
  const navigate = useNavigate();

  const onSuccess = async res => {
    try {
      const response = await googleSignIn(res);
      console.log(response);

      if (response.status === 200) {
        localStorage.setItem('user', JSON.stringify(response.data));
        navigate('/signup/policy');
      } else if (response.status === 201) {
        navigate('/home');
      }
    } catch (e) {
      console.log(e);
      alert(e);
    }
  };

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
        <Header title="Login" isBack />
        <Flex
          flexDirection="column"
          alignItems="center"
          p="40px 10px 10px"
          gap="24px"
        >
          <Text fontSize="3xl">Login to Bard</Text>
          <Image boxSize="80px" src={ViolinIcon} />
          <GoogleLogin
            onSuccess={onSuccess}
            onError={res => console.log(res)}
          />
        </Flex>
      </Flex>
    </>
  );
}

export default Login;
