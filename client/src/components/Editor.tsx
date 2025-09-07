
import { useQuill } from 'react-quilljs';

import 'quill/dist/quill.snow.css'; // Add css for snow theme
import { useEffect } from 'react';
// or import 'quill/dist/quill.bubble.css'; // Add css for bubble theme
const Editor = ({setText, initialText}: {initialText: string, setText: React.Dispatch<React.SetStateAction<string>>}) => {
    
    const { quill, quillRef } = useQuill({
        theme: 'snow',
        modules: {
            toolbar: [
                ['bold', 'italic', 'underline'],
            ]
        },
        placeholder: 'Type here',
        formats: [
            'header',
            'size',
            'bold',
            'italic',
            'underline',
        ]
    })

    useEffect(() => {
        if (quill) {
          quill.root.innerHTML = initialText;
          quill.on('text-change', () => {
            setText(quill.root.innerHTML);
          });
        }
      }, [quill, setText, initialText]);
    
  return (
    <div>
        <div ref={quillRef} />

    </div>
  )
}

export default Editor