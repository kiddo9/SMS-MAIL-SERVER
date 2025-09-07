import { Loader2, Pencil, PlusCircle, Trash2 } from "lucide-react"
import { useEffect, useState } from "react";
import CreateTemplate from "../components/CreateTemplate";
import { toast } from "react-toastify";
import TemplateClient from "../lib/templateClient";
import type { SmsTemplate, Template} from "../proto/Template";
import DeleteTemplate from "../components/DeleteTemplate";
import EditTemplate from "../components/EditTemplate";

const Templates = () => {
  const [openCreate, setOpenCreate] = useState(false);
  const [openEdit, setOpenEdit] = useState(false);
  const [openDelete, setOpenDelete] = useState(false);
  const [deleteId, setDeleteId] = useState('');
  const [deleteName, setDeleteName] = useState('');
  const [deleteType, setDeleteType] = useState<'email' | 'sms'>();
  const [editId, setEditId] = useState('');
  const [editType, setEditType] = useState<'email' | 'sms'>();
  const [loading , setLoading] = useState(false);
  const [smsTemplates, setSmsTemplates] = useState<SmsTemplate[]>([]);
  const [emailTemplates, setEmailTemplates] = useState<Template[]>([]);
  const [reload, setReload] = useState(false);

  useEffect(() => {
    const getTemplates = async () => {
      setLoading(true);
      try {
        const emailArray: Template[] = [];
        const smsArray: SmsTemplate[] = [];
        const tempStream = TemplateClient.allTemplates({});
        for await (const response of tempStream.responses) {
          
          if(response.emailTemplate){
            console.log(response);
            emailArray.push(response.emailTemplate);
            setEmailTemplates(emailArray);
          }
          if(response.smsTemplate){
            console.log(response);
            smsArray.push(response.smsTemplate);
            setSmsTemplates(smsArray);
          }
        }
      } catch (error) {
        toast.error(error instanceof Error ? error.message : "An unexpected error occurred.");
        if(import.meta.env.VITE_ENV === "development") console.error(error);
      }finally{
        setLoading(false);
        setReload(false);
      }
    }
    getTemplates();
  }, [reload]);
  return (
    <div className="min-h-[calc(100vh-120px)] bg-gradient-to-br from-blue-50 via-white to-indigo-50">
        <div className="flex flex-col mx-auto px-6 py-8">
          <div className="flex relative w-full justify-center items-center px-4 py-2 rounded-full text-lg font-bold mb-6 text-center">
            <img className="rotate-y-180" width={30} src="/logo-icon.png" alt="Logo"/>
            <span>Templates</span>
            <img className="" width={30} src="/logo-icon.png" alt="Logo" />
            <button onClick={() => setOpenCreate(true)} className="flex hover:bg-blue-500 transition-all absolute right-0 cursor-pointer bg-[#6699ff] items-center gap-2 px-4 py-2 rounded-full text-lg font-bold text-center">
                <PlusCircle size={20} />
                <span>Add</span>
            </button>
          </div>
        </div>
        <div className='py-2 border-t-2 border-gray-300 bg-gray-100 grid grid-cols-5 lg:grid-cols-7 gap-10 w-full items-center justify-items-center px-5 shadow-md'>
            <h1 className='text-sm  text-black/50 font-semibold justify-self-start'>S/N</h1>
            <h1 className='text-sm  text-black/50 font-semibold justify-self-start'>TYPE</h1>
            <h1 className='text-sm  text-black/50 font-semibold justify-self-start'>NAME</h1>
            <h1 className='text-sm hidden lg:block  text-black/50 font-semibold col-span-2'>CONTENT</h1>
            <h1 className='text-sm  text-black/50 font-semibold justify-self-center'>CREATED</h1>
            <h1 className='text-sm  text-black/50 font-semibold justify-self-end'>ACTIONS</h1>
        </div>
        {
          loading 
          ?
            <Loader2 size={50} className="mx-auto mt-10 animate-spin" />
          : 
            (() => {
              if(smsTemplates.length === 0 && emailTemplates.length === 0) return <p className="text-center text-sm text-gray-400 mt-10">No templates created yet.</p>
              return (
                <>
                  {smsTemplates.map((sms, index) => (
                    <div key={index} className='py-2 grid grid-cols-5 lg:grid-cols-7 gap-10 w-full items-center justify-items-center px-5 border-b-2 border-gray-300'>
                        <p className='text-sm  text-black/50 font-semibold justify-self-start'>{index + 1}</p>
                        <p className='text-sm  text-black/50 font-semibold justify-self-start'>SMS</p>
                        <p className='text-sm  text-black/50 font-semibold justify-self-start'>{sms.smsTemplateName}</p>
                        <p className='text-sm  hidden lg:block text-black/50 col-span-2 font-semibold lg:max-w-[200px] xl:max-w-[350px] justify-self-start text-nowrap text-ellipsis overflow-hidden'>{sms.smsTemplateContent}</p>
                        <p className='text-sm  text-black/50 font-semibold justify-self-center'>{new Date(sms.date).toLocaleString()}</p>
                        <div className="flex justify-self-end py-2 gap-2">
                            <button onClick={() => {
                                setOpenEdit(true);
                                setEditId(sms.id);
                                setEditType('sms');
                            }} className=" hover:text-blue-500 transition-all cursor-pointer items-center rounded-full text-lg font-bold text-center">
                                <Pencil size={20} />
                            </button>
                            <button onClick={() => {
                                setOpenDelete(true);
                                setDeleteId(sms.id);
                                setDeleteName(sms.smsTemplateName);
                                setDeleteType('sms');
                            }} className=" hover:text-red-500 transition-all cursor-pointer rounded-full text-lg font-bold text-center">
                                <Trash2 size={20} />
                            </button>
                        </div>
                    </div>
                  ))}
                  {emailTemplates.map((email, index) => (
                    <div key={index} className='py-2 grid grid-cols-5 lg:grid-cols-7 gap-10 w-full items-center justify-items-center px-5 border-b-2 border-gray-300'>
                        <p className='text-sm  text-black/50 font-semibold justify-self-start'>{smsTemplates.length + index + 1}</p>
                        <p className='text-sm  text-black/50 font-semibold justify-self-start'>EMAIL</p>
                        <p className='text-sm  text-black/50 font-semibold justify-self-start'>{email.templateName}</p>
                        <p className='text-sm hidden lg:block  text-black/50 col-span-2 font-semibold lg:max-w-[200px] xl:max-w-[350px] justify-self-start text-nowrap text-ellipsis overflow-hidden'>{email.templateContent}</p>
                        <p className='text-sm  text-black/50 font-semibold justify-self-center'>{new Date(email.date).toLocaleString()}</p>
                        <div className="flex justify-self-end py-2 gap-2">
                            <button onClick={() => {
                                setEditType("email");
                                setEditId(email.id);
                                setOpenEdit(true);
                            }} className=" hover:text-blue-500 transition-all cursor-pointer items-center rounded-full text-lg font-bold text-center">
                                <Pencil size={20} />
                            </button>
                            <button onClick={() => {
                                setDeleteType("email");
                                setDeleteId(email.id);
                                setDeleteName(email.templateName);
                                setOpenDelete(true);
                            }} className=" hover:text-red-500 transition-all cursor-pointer rounded-full text-lg font-bold text-center">
                                <Trash2 size={20} />
                            </button>
                        </div>
                    </div>
                  ))}
                </>
              )
            })()
        }
        {openCreate && <CreateTemplate setReload={setReload} setOpenCreate={setOpenCreate} />}
        {openDelete && <DeleteTemplate setReload={setReload} type={deleteType} id={deleteId} setOpenDelete={setOpenDelete} name={deleteName} />}
        {openEdit && <EditTemplate setReload={setReload} type={editType} setOpenEdit={setOpenEdit} id={editId}/>}
    </div>
  )
}

export default Templates