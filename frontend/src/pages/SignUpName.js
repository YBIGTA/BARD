import { Button, Text, Flex, Input } from '@chakra-ui/react';
import { useEffect, useState } from 'react';
import { useNavigate } from 'react-router';
import { signupUser } from '../apis/user';
import BottomButton from '../components/BottomButton';
import Header from '../components/Header';

function SignUpName() {
  const navigate = useNavigate();

  const [userInfo, setUserInfo] = useState(
    JSON.parse(localStorage.getItem('user'))
  );

  const [username, setUsername] = useState(userInfo.name);
  const [showError, setShowError] = useState(false);

  const onSubmit = async e => {
    e.preventDefault();
    if (username.trim() === '') {
      setShowError(true);
      return;
    }
    const res = await signupUser({
      email: userInfo.email,
      name: username,
      social_id: userInfo.social_id,
    });

    if (res.status === 201) {
      alert('Sign up success!');
      navigate('/login');
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
        <Header title="Sign Up" isBack />
        <Flex
          flexDirection="column"
          justifyContent="space-between"
          align="center"
          p="60px 24px"
        >
          <Flex
            flexDirection="column"
            width="100%"
            align="flex-start"
            gap="0.5rem"
          >
            <Input
              type="text"
              value={username}
              onChange={e => {
                setShowError(false);
                setUsername(e.target.value);
              }}
              placeholder="Username"
              size="lg"
            />
            {showError && (
              <Text fontSize="xs" color="red.500">
                Username is required
              </Text>
            )}
          </Flex>
          <BottomButton onClick={onSubmit} title="Submit" />
        </Flex>
      </Flex>
    </>
  );
}

export default SignUpName;
