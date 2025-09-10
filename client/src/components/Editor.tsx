
// import EditorJS from '@editorjs/editorjs';

// import { useEffect, useRef } from 'react';
// import Header from "@editorjs/header";
// import List from "@editorjs/list";
// // or import 'quill/dist/quill.bubble.css'; // Add css for bubble theme
// const Editor = ({setText, initialText}: {initialText: string, setText: React.Dispatch<React.SetStateAction<string>>}) => {
    
//     const editorRef = useRef<EditorJS>(null);

//   useEffect(() => {
//     if (!editorRef.current) {
//       const editor = new EditorJS({
//         holder: "editor",
//         autofocus: true,
//         tools: {
//           header: Header,
//           list: List,
//         },
//         data: "",
//         onReady: () => {
//           editorRef.current = editor;
//         },

//         onChange: async () => {
//           const content = await editor.save();
//           console.log(content);
//         }
//       });
//     }

//   }, [initialText, setText]);

    
    
//   return (
//     <div>
//         <div id="editor"></div>

//     </div>
//   )
// }

// export default Editor