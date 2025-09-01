
import { useQuill } from 'react-quilljs';

import 'quill/dist/quill.snow.css'; // Add css for snow theme
import { useEffect } from 'react';
// or import 'quill/dist/quill.bubble.css'; // Add css for bubble theme
const Editor = ({setText}: {text: string, setText: React.Dispatch<React.SetStateAction<string>>}) => {
    
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
          quill.root.innerHTML = `
Hi {{.Name}},

This is a friendly reminder about your outstanding balance of ₦{{.PendingPrice}} for your {{.course}} course. To ensure 
continued access to your lectures, please settle your balance by or before {{.Date}}.<br/>
For your convienience, you can make your payment through online bank transfer. <br/>

if you've already made the payment, please meet with the admin with prove of payment Otherwise, you can contact the Admin department at {{.phoneNumber}} or 
{{.EmailAddress}} to discuss payment options.

Thanks,
Neo cloud Admin Department.
`;
          quill.on('text-change', () => {
            setText(quill.root.innerHTML);
          });
        }
      }, [quill, setText]);
    
  return (
    <div>
        <div ref={quillRef} />

    </div>
  )
}

export default Editor