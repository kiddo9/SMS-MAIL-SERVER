import { useState } from "react";
import AdminClient from "../../lib/adminClient"
import {OtpRequest} from "../../proto/Admin"

const Login = () => {
    const [email, setEmail] = useState('');
    // const [password, setPassword] = useState('');
    const handleSubmit = async(e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();

        const response = await AdminClient.loginAdmin(OtpRequest.create({email: email}), {
           "recaptcha-token": "recaptcha-token" 
        })

        
        console.log(response);
    }
  return (
    <div className='flex flex-col h-screen justify-center items-center bg-[#6699ff]/10'>
        
        <form onSubmit={handleSubmit} className='bg-white p-4 rounded-md shadow-md flex flex-col min-w-[300px] max-w-[400px]'>
            <img className='self-center' width={100} src="./logo-icon.png" alt="Logo" />
            <h1 className='text-2xl font-bold mb-4 text-center'>Sign In</h1>
            <div className="mb-4">
                <label htmlFor="email" className='block mb-1 text-sm'>Email<span className='text-red-500'>*</span></label>
                <input value={email} onChange={(e) => setEmail(e.target.value)} className='p-2 border-2 focus:border-[#6699ff] border-gray-200 rounded-md w-full' type="text" placeholder='Email' />
            </div>
            {/* <div>
                <label htmlFor="password" className='block mb-1 text-sm'>Password</label>
                <input value={password} onChange={(e) => setPassword(e.target.value)} className='p-2 border-2 focus:border-[#6699ff] border-gray-200 rounded-md w-full' type="password" placeholder='Password' />
            </div> */}
            <button type="submit" className='mt-4 bg-[#6699ff] hover:bg-[#6699ff]/80 cursor-pointer text-white p-2 rounded-md'>Sign In</button>
        </form>
    </div>

  )
}

export default Login