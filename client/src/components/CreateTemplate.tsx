import { useState } from "react"
import Editor from "./Editor"
import TemplateClient from "../lib/templateClient";
import { toast } from "react-toastify";
import { demoTemplate } from "../lib/constants";

const CreateTemplate = ({setOpenCreate, setReload}: {setOpenCreate: React.Dispatch<React.SetStateAction<boolean>>, setReload: React.Dispatch<React.SetStateAction<boolean>>}) => {
    const [loader, setLoader] = useState(false)
    const [name, setName] = useState('')
    const [type, setType] = useState<"email" | "sms">("sms")
    const [text, setText] = useState('')

    const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
        setLoader(true)
        e.preventDefault();
        {/*CREATE API GOES HERE */}
        switch (type) {
            case "email":
                try {
                    const request = await TemplateClient.createEmailTemplate({
                        date: new Date(Date.now()).toString(),
                        templateContent: text,
                        templateName: name
                    })
                    const response = request.response;
                    if(response.status == true){
                        toast.success(response.message);
                        setOpenCreate(false);
                        setReload(true);
                        return
                    }
                    toast.error(response.message);
                } catch (error) {
                    toast.error(error instanceof Error ? error.message : "An unexpected error occurred.");
                    if(import.meta.env.VITE_ENV === "development") console.error(error);
                    
                }finally{
                    setLoader(false)
                }
                break;
            case "sms":
                try {
                    const request = await TemplateClient.createSmsTemplate({
                        date: new Date(Date.now()).toString(),
                        smsTemplateContent: text,
                        smsTemplateName: name
                    })
                    const response = request.response;
                    if(response.status == true){
                        toast.success(response.message);
                        setOpenCreate(false);
                        setReload(true);
                        return
                    }
                    toast.error(response.message);
                } catch (error) {
                    toast.error(error instanceof Error ? error.message : "An unexpected error occurred.");
                    if(import.meta.env.VITE_ENV === "development") console.error(error);
                    
                }finally{
                    setLoader(false)
                }
                break;
            default:
                setLoader(false)
                toast.error("Invalid template type");
                break;
        }
        
    }
    
  return (
    <div   className='absolute top-0 left-0 flex justify-center items-center backdrop-blur-xs w-full h-full z-50 '>
        <div onClick={() => setOpenCreate(false)}  className='absolute top-0 left-0 w-full h-full bg-black opacity-50'/>
        <div className='bg-white py-2 rounded-lg shadow-2xl z-10 w-[600px]'>
            <div className="flex relative w-full justify-center items-center px-4 py-2 rounded-full text-lg font-bold mb-2 text-center">
                <img className="rotate-y-180" width={30} src="/logo-icon.png" alt="Logo"/>
                <span>Create Template</span>
                <img className="" width={30} src="/logo-icon.png" alt="Logo" />
            </div>
            <div className='px-5 max-h-[40vh] overflow-y-scroll'>
                <form onSubmit={handleSubmit} className='flex flex-col'>
                    <fieldset className="mb-4 flex flex-col gap-1">
                        <label className="text-sm" htmlFor="name">
                            Name
                        </label>
                        <input
                            onChange={(e) => setName(e.target.value)}
                            value={name}
                            className="rounded-lg px-4 py-2 outline-none border-2 border-gray-500 focus:border-[#6699ff]"
                            type="text"
                            id="name"
                            name="name"
                            placeholder="Enter template name"
                            required
                        />
                    </fieldset>
                    <fieldset className="mb-4 flex flex-col gap-1">
                        <label className="text-sm" htmlFor="type">
                            Type
                        </label>
                        <select
                            onChange={(e) => setType(e.target.value as "email" | "sms")}
                            value={type}
                            className="rounded-lg appearance-none px-4 py-2 outline-none border-2 border-gray-500 focus:border-[#6699ff]"
                            id="type"
                            name="type"
                            required
                        >
                            <option  value="email">Email</option>
                            <option value="sms">SMS</option>
                        </select>
                    </fieldset>
                    <p className="text-sm mb-4"><strong>Instructions: </strong>{"You can edit the text below to create a custom template, please do not edit or remove the text in '{ {. } }' symbol"}</p>
                    <Editor initialText={demoTemplate} setText={setText} />
                    <button
                        type="submit"
                        className="border-2 mt-2 border-[#6699ff] mb-5 mx-auto text-[#6699ff] hover:bg-blue-500 hover:text-white px-4 py-2 rounded-lg transition duration-300 ease-in cursor-pointer hover:shadow-2xl"
                    >
                        {loader ? 'Creating...' : 'Create'}
                    </button>
                </form>
            </div>
            

       </div>
    </div>
  )
}

export default CreateTemplate  