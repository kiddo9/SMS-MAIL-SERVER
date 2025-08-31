import { PlusCircle } from "lucide-react"
import { useState } from "react";
import CreateTemplate from "../components/CreateTemplate";

const Templates = () => {
  const [openCreate, setOpenCreate] = useState(true);
  const [openEdit, setOpenEdit] = useState(false);
  const [openDelete, setOpenDelete] = useState(false);
  return (
    <div className="min-h-screen bg-gradient-to-br from-blue-50 via-white to-indigo-50">
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
        <div className='py-2 border-t-2 border-gray-300 bg-gray-100 grid grid-cols-6 gap-10 w-full items-center justify-items-center px-5 shadow-md'>
            <h1 className='text-sm  text-black/50 font-semibold justify-self-start'>S/N</h1>
            <h1 className='text-sm  text-black/50 font-semibold justify-self-start'>NAME</h1>
            <h1 className='text-sm  text-black/50 font-semibold'>DESCRIPTION</h1>
            <h1 className='text-sm  text-black/50 font-semibold justify-self-center'>CREATED</h1>
            <h1 className='text-sm  text-black/50 font-semibold justify-self-end'>ACTIONS</h1>
        </div>
        {openCreate && <CreateTemplate setOpenCreate={setOpenCreate} />}
        {openEdit && <CreateTemplate setOpenCreate={setOpenEdit} />}
        {openDelete && <CreateTemplate setOpenCreate={setOpenDelete} />}
    </div>
  )
}

export default Templates