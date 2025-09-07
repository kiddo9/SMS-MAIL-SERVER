import React from 'react'
import { toast } from 'react-toastify';


const DeleteTemplate = ({setOpenDelete, type, name, id}: {setOpenDelete: React.Dispatch<React.SetStateAction<boolean>>, type: 'email' | 'sms' | undefined, name: string, id: string}) => {
    const [loader, setLoader] = React.useState(false);
    const handleDelete = async() => {
        setLoader(true);
        try {
            switch (type) {
                case "email":
                    try {
                        //api goes here
                    } catch (error) {
                        toast.error(error instanceof Error ? error.message : "An unexpected error occurred.");
                        if(import.meta.env.VITE_ENV === "development") console.error(error);
                        
                    }
                    break;
                case "sms":
                    try {
                        //api goes here
                    } catch (error) {
                        toast.error(error instanceof Error ? error.message : "An unexpected error occurred.");
                        if(import.meta.env.VITE_ENV === "development") console.error(error);
                    }
                    break;
                default:
                    break;
            }
        } catch (error) {
            toast.error(error instanceof Error ? error.message : "An unexpected error occurred.");
            if(import.meta.env.VITE_ENV === "development") console.error(error);
        }finally{
            setLoader(false);
        }
        
    }
  return (
    <div   className='absolute top-0 left-0 flex justify-center items-center backdrop-blur-xs w-full h-full z-50 '>
        <div onClick={() => setOpenDelete(false)}  className='absolute top-0 left-0 w-full h-full bg-black opacity-50'/>
       <div className='bg-white py-2 rounded-lg shadow-2xl z-10 w-[400px]'>
            <header className="bg-[#D7DDFF] w-full flex flex-row items-center px-4 py-2 shadow-md">
                <h1 className="text-xl mx-auto text-red-500">Delete {name}</h1>
            </header>
            <p className='text-center mt-5 mb-2'>
                {`Are you sure you want to delete this ${type}?`}
                <br />
                <span className='font-bold'>{name}</span>
            </p>
            <div className='flex justify-center gap-4 px-10 pb-3'>
                <button disabled={loader} onClick={handleDelete} className='bg-red-500 text-white px-4 py-2 rounded-lg cursor-pointer hover:bg-red-600'>
                    {loader ? 'Deleting...' : 'Yes'}
                </button>
                <button disabled={loader} onClick={() => setOpenDelete(false)} className='bg-[#6699ff] text-white px-4 py-2 rounded-lg cursor-pointer hover:shadow-2xl'>No</button>
            </div>
       </div>
    </div>
  )
}

export default DeleteTemplate