import { useEffect, useState } from "react";
import SMSAPIClient from '../lib/smsApiClient';
import { toast } from "react-toastify";
import CountUp from "react-countup";



const SMSBalances = () => {
    const [data, setData] = useState({
        ebulk: '',
        bulk: ''
    });

    useEffect(() => {
        const fetchBalances = async () => {
            try {
                const ebulkStream = SMSAPIClient.ebulkSmsWallet({});
                for await (const response of ebulkStream.responses) {
                    // console.log(response);
                    if(isNaN(Number(response.response))){
                        toast.warning("EBulk SMS wallet balance is not available. Please notify the service administrator.");
                    };
                    setData(prevData => ({
                        ...prevData,
                        ebulk: response.response
                    }));
                }
                const bulkStream = SMSAPIClient.bulkSmsWallet({});
                for await (const response of bulkStream.responses) {
                    // console.log(response);
                    if(isNaN(Number(response.response))){
                        toast.warning("Bulk SMS NG wallet balance is not available. Please notify the service administrator.", {
                            
                        });
                    };
                    setData(prevData => ({
                        ...prevData,
                        bulk: response.response
                    }));
                }
            } catch (error) {
                toast.error(error instanceof Error ? error.message : "An unexpected error occurred.");
                if(import.meta.env.VITE_ENV === "development") console.error(error);
            }finally{
                // setLoading(false);
            }
            
            
        }

        fetchBalances();
        
    }, [])
  return (
    <div className="min-h-[calc(100vh-120px)] bg-gradient-to-br from-blue-50 via-white to-indigo-50">
        <div className="flex flex-col mx-auto px-6 py-8">
            <div className="flex relative w-full justify-center items-center px-4 py-2 rounded-full text-lg font-bold mb-6 text-center">
                <img className="rotate-y-180" width={30} src="/logo-icon.png" alt="Logo"/>
                <span>SMS Wallet Balances</span>
                <img className="" width={30} src="/logo-icon.png" alt="Logo" />
            </div>
            
            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
                <div className="bg-white rounded-lg shadow-lg p-6 max-w-sm">
                    <div className="flex items-center justify-between mb-3">
                        <h3 className="text-lg font-semibold text-gray-800">Ebulk SMS</h3>
                        <span className="text-sm text-green-600 bg-green-100 px-2 py-1 rounded">Active</span>
                    </div>
                        
                    <div className="text-3xl font-bold text-gray-900 mb-2">
                        ₦{Number(data.ebulk) 
                            ? <CountUp decimal="." separator="," decimals={2} end={Number(data.ebulk)} duration={2} /> 
                            : "0.00"
                        }
                    </div>
                    <p className="text-gray-500 text-sm">You can top up this wallet at <a className="text-[#6699ff]" href="https://www.ebulksms.com/login" target="_blank">EBulk SMS</a></p>
                </div>

                <div className="bg-white rounded-lg shadow-lg p-6 max-w-sm">
                    <div className="flex items-center justify-between mb-3">
                        <h3 className="text-lg font-semibold text-gray-800">Bulk SMS NG</h3>
                        <span className="text-sm text-green-600 bg-green-100 px-2 py-1 rounded">Active</span>
                    </div>
                        
                    <div className="text-3xl font-bold text-gray-900 mb-2">
                        ₦{Number(data.bulk) 
                            ? <CountUp decimal="." separator="," decimals={2} end={Number(data.bulk)} duration={2} /> 
                            : "0.00" 
                        }
                    </div>
                    <p className="text-gray-500 text-sm">You can top up this wallet at <a className="text-[#6699ff]" href="https://www.bulksmsnigeria.com/login" target="_blank">Bulk SMS NG</a></p>
                </div>

            </div>
            
        </div>
    </div>
  )
}

export default SMSBalances