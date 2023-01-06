import { Button, Checkbox, Divider, Flex, Stack } from '@chakra-ui/react';
import { useState } from 'react';
import { useNavigate } from 'react-router';
import BottomButton from '../components/BottomButton';
import Header from '../components/Header';

function SignUpPolicy() {
  const [checkedItems, setCheckedItems] = useState([false, false]);
  const allChecked = checkedItems.every(Boolean);
  const isIndeterminate = checkedItems.some(Boolean) && !allChecked;

  const navigate = useNavigate();
  const onClick = e => {
    e.preventDefault();
    if (!allChecked) {
      alert('Please agree to all policies.');
      return;
    }
    navigate('/signup/user');
  };

  return (
    <Flex
      w="100%"
      flexDirection="column"
      flexGrow="1"
      flexShrink="1"
      flexBasis="0%"
      overflow="auto"
    >
      <Header title="Sign up" isBack />
      <Flex
        flexDirection="column"
        justifyContent="space-between"
        align="center"
        p="60px 24px"
      >
        <Flex
          flexDirection="column"
          align="flex-start"
          p="0px"
          gap="16px"
          w="calc(100vw - 48px)"
          maxW="calc(768px - 48px)"
        >
          <Checkbox
            colorScheme="yellow"
            isChecked={allChecked}
            isIndeterminate={isIndeterminate}
            onChange={e => {
              e.preventDefault();
              setCheckedItems(prev => prev.map(checked => !checked));
            }}
          >
            Agree all
          </Checkbox>
          <Divider />
          <Stack spacing={1}>
            <Checkbox
              colorScheme="yellow"
              isChecked={checkedItems[0]}
              onChange={e => {
                e.preventDefault();
                setCheckedItems(prev => [!prev[0], prev[1]]);
              }}
            >
              I agree to terms of service
            </Checkbox>
            <Checkbox
              colorScheme="yellow"
              isChecked={checkedItems[1]}
              onChange={e => {
                e.preventDefault();
                setCheckedItems(prev => [prev[0], !prev[1]]);
              }}
            >
              I agree to Privacy Policy
            </Checkbox>
          </Stack>
        </Flex>
        <BottomButton title="submit" type="button" onClick={onClick} />
      </Flex>
    </Flex>
  );
}

export default SignUpPolicy;
