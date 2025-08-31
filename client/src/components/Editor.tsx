
import { useQuill } from 'react-quilljs';

import 'quill/dist/quill.snow.css'; // Add css for snow theme
import { useEffect } from 'react';
// or import 'quill/dist/quill.bubble.css'; // Add css for bubble theme
const Editor = () => {

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
          quill.on('text-change', () => {
            console.log(quill.getText());
          });
        }
      }, [quill]);
    
  return (
    <div>
        <div ref={quillRef} />

    </div>
  )
}

export default Editor