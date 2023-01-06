import { useCallback, useEffect, useMemo, useRef } from 'react';

const useAppearSentencesOnScroll = () => {
  const sentenceRefs = useRef([]);

  const observer = useMemo(
    () =>
      new IntersectionObserver(entries => {
        entries.forEach(entry => {
          console.log(entry);
          console.log(entry.target.classList);
          if (entry.isIntersecting) {
            entry.target.style.opacity = 1;
            entry.target.style.filter = 'blur(0)';
            entry.target.style.transform = 'translateX(0)';
          } else {
            entry.target.style.filter = 'blur(5px)';
            entry.target.style.opacity = 0;
            entry.target.style.transition = 'all 1s';
          }
        });
      }),
    []
  );

  const observeSentences = useCallback(() => {
    sentenceRefs && sentenceRefs.current.forEach(el => observer.observe(el));
  }, [observer]);

  useEffect(() => {
    window.addEventListener('scroll', observeSentences);

    return () => {
      window.removeEventListener('scroll', observeSentences);
    };
  }, [observeSentences]);

  return { sentenceRefs };
};

export default useAppearSentencesOnScroll;
