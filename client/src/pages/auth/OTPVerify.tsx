import { useEffect, useState } from "react";
import AdminClient from "../../lib/adminClient"
import {OtpRequest, OtpVerificationRequest} from "../../proto/Admin"

const OTPVerify = () => {
    const [otp, setOtp] = useState('');
    const [allowReset, setAllowReset] = useState(false);
    
    const email = '';
    // const [password, setPassword] = useState('');
    const handleSubmit = async(e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();

        const response = await AdminClient.verifyOtp(OtpVerificationRequest.create({email: email, otp: otp}),{})

        
        console.log(response);
    }

    const handleOTPResend = async() =>{
        setAllowReset(false);
        const response = await AdminClient.sendOtp(OtpRequest.create({email: email}),{})
        console.log(response);
    }

    useEffect(() => {
        if(!allowReset){
            const timer = setTimeout(() => {
                setAllowReset(true);
            }, 5000);
            return () => clearTimeout(timer);
        }
        else return
    }, [allowReset])
  return (
    <div className='flex flex-col h-screen justify-center items-center bg-[#6699ff]/10'>
        
        <form onSubmit={handleSubmit} className='bg-white p-4 rounded-md shadow-md flex flex-col min-w-[300px] max-w-[400px]'>
            <div className="flex flex-row-reverse  items-center justify-center">
                <img className='' width={30} src="/logo-icon.png" alt="Logo" />
                <h1 className='text-lg font-bold mb-1 text-center'>Verify OTP</h1>
                <img className='rotate-y-180' width={30} src="/logo-icon.png" alt="Logo" />
            </div>
            <span className='block mb-1 text-xs text-center text-gray-500 '>An OTP has been sent to your email</span>
            <div className="mt-2 mb-1">
                
                <input value={otp} onChange={(e) => setOtp(e.target.value)} className='p-2 border-2 focus:border-[#6699ff] border-gray-200 rounded-md w-full' type="text" placeholder='Enter the OTP...' />
            </div>
            <button type="submit" className='mt-1 bg-[#6699ff] hover:bg-[#6699ff]/80 cursor-pointer text-white p-2 rounded-md'>Verify</button>
            <button disabled={!allowReset} onClick={handleOTPResend} className='mt-1 bg-transparent self-start text-xs font-semibold cursor-pointer text-[#6699ff] disabled:text-gray-500 disabled:cursor-not-allowed'>Resend</button>
        </form>
    </div>

  )
}

export default OTPVerify