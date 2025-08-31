import { useState, useCallback } from 'react';
import { Upload, FileSpreadsheet, MessageSquare, CheckCircle, AlertCircle, X, NotepadText} from 'lucide-react';
import { Link } from 'react-router-dom';
import FileUploadClient from '../lib/fileUploadClient';
import {fileUploadRequest} from '../proto/FileUpload'

const Home = () => {
  const [dragActive, setDragActive] = useState(false);
  const [file, setFile] = useState<File | null>(null);
  const [uploading, setUploading] = useState(false);
  const [uploadStatus, setUploadStatus] = useState<string | null>(null); // 'success', 'error', null

  const handleDrag = useCallback((e: React.DragEvent<HTMLDivElement>) => {
    e.preventDefault();
    e.stopPropagation();
    if (e.type === "dragenter" || e.type === "dragover") {
      setDragActive(true);
    } else if (e.type === "dragleave") {
      setDragActive(false);
    }
  }, []);

  const handleDrop = useCallback((e: React.DragEvent<HTMLDivElement>) => {
    e.preventDefault();
    e.stopPropagation();
    setDragActive(false);
    
    const files = e.dataTransfer?.files;
    if (files && files[0]) {
      handleFileSelect(files[0]);
    }
  }, []);

  const handleFileSelect = (selectedFile: File) => {
    const allowedTypes = [
      'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet',
      'text/csv',
      'application/vnd.ms-excel'
    ];
    
    if (allowedTypes.includes(selectedFile.type) || 
        selectedFile.name.endsWith('.xlsx') || 
        selectedFile.name.endsWith('.csv')) {
      setFile(selectedFile);
      setUploadStatus(null);
    } else {
      setUploadStatus('error');
    }
  };

  const handleFileInput = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (e.target.files && e.target.files[0]) {
      handleFileSelect(e.target.files[0]);
    }
  };

  const handleUpload = async () => {
    if (!file) return;
    
    setUploading(true);
    
    // Simulate file processing
    try {
      const reader = new FileReader();
      reader.onload = async(event) => {
        const arrayBuffer = event.target?.result as ArrayBuffer;   // raw file bytes
        const uint8Array = new Uint8Array(arrayBuffer);
        const request = await FileUploadClient.fileUpload(fileUploadRequest.create({
          content: uint8Array,
          date: new Date(Date.now() + 24 * 60 * 60 * 1000).toISOString()
        }), {
          meta: { 'send_using': 'EBulksms' } // Example metadata
        });

        const response = request.response;

        console.log(response);
      }
      reader.readAsArrayBuffer(file);
      setUploadStatus('success');
    } catch (error) {
      if(import.meta.env.VITE_ENV === "development") console.error(error);
      setUploadStatus('error');
    } finally {
      setUploading(false);
    }
  };

  const resetUpload = () => {
    setFile(null);
    setUploadStatus(null);
    setUploading(false);
  };

  return (
    <div className="min-h-screen bg-gradient-to-br from-blue-50 via-white to-indigo-50">
      <div className="max-w-4xl flex flex-col mx-auto px-6 py-12">
        {/* Hero Section */}
        <div className="text-center mb-4">
          <div className="flex justify-center items-center px-4 py-2 rounded-full text-lg font-bold mb-6 text-center">
            <img className="rotate-y-180" width={30} src="/logo-icon.png" alt="Logo"/>
            <span>SMS Notifications Made Easy</span>
            <img className="" width={30} src="/logo-icon.png" alt="Logo" />
          </div>
          
          <h2 className="text-4xl font-bold text-gray-900 mb-4">
            Automated Fee Reminder System
          </h2>
          
          <p className="text-xl text-gray-600 mb-8 max-w-2xl mx-auto">
            Upload Neo Cloud student data file and automatically send personalized SMS reminders to students with outstanding fees. Simple, efficient, and effective.
          </p>

          {/* Feature Cards */}
          <div className="grid md:grid-cols-3 gap-6 mb-12">
            <div className="bg-white rounded-xl p-6 shadow-sm border border-gray-100 hover:shadow-md transition-shadow">
              <FileSpreadsheet className="w-10 h-10 text-[#6699ff] mx-auto mb-4" />
              <h3 className="font-semibold text-gray-800 mb-2">Easy Upload</h3>
              <p className="text-sm text-gray-600">Support for Excel (.xlsx) and CSV files with student data</p>
            </div>
            
            <div className="bg-white rounded-xl p-6 shadow-sm border border-gray-100 hover:shadow-md transition-shadow">
              <NotepadText className="w-10 h-10 text-green-600 mx-auto mb-4" />
              <h3 className="font-semibold text-gray-800 mb-2">Templating</h3>
              <p className="text-sm text-gray-600">Create personalized SMS templates</p>
            </div>
            
            <div className="bg-white rounded-xl p-6 shadow-sm border border-gray-100 hover:shadow-md transition-shadow">
              <MessageSquare className="w-10 h-10 text-purple-600 mx-auto mb-4" />
              <h3 className="font-semibold text-gray-800 mb-2">Bulk SMS</h3>
              <p className="text-sm text-gray-600">Send reminders to multiple students instantly</p>
            </div>
          </div>
        </div>

        <Link to="/templates" className='self-center cursor-pointer'>
          <button className="mb-12 px-8 py-3  bg-transparent border-2 border-[#6699ff] text-[#6699ff] font-semibold rounded-lg hover:bg-blue-500 hover:text-white disabled:opacity-50 disabled:cursor-not-allowed cursor-pointer transition-colors shadow-xs">
            Manage Templates
          </button>
        </Link>
        

        {/* Upload Section */}
        <div className="bg-white rounded-2xl shadow-xl border border-gray-200 p-8">
          <div className="text-center mb-6">
            <h3 className="text-2xl font-bold text-gray-800 mb-2">Upload Student Data</h3>
            <p className="text-gray-600">Upload your Excel or CSV file containing student names, phone numbers, and fee balances</p>
          </div>

          {!file ? (
            <div
              className={`relative border-2 border-dashed rounded-xl p-12 text-center transition-all duration-300 ${
                dragActive 
                  ? 'border-blue-500 bg-blue-50' 
                  : 'border-gray-300 bg-gray-50 hover:border-blue-400 hover:bg-blue-25'
              }`}
              onDragEnter={handleDrag}
              onDragLeave={handleDrag}
              onDragOver={handleDrag}
              onDrop={handleDrop}
            >
              <Upload className={`w-16 h-16 mx-auto mb-6 ${dragActive ? 'text-[#6699ff]' : 'text-gray-400'}`} />
              
              <h4 className="text-xl font-semibold text-gray-700 mb-2">
                {dragActive ? 'Drop your file here' : 'Drag and drop your file here'}
              </h4>
              
              <p className="text-gray-500 mb-6">
                or click to browse your files
              </p>

              <label className="inline-flex items-center px-6 py-3 bg-[#6699ff] text-white font-medium rounded-lg hover:bg-blue-700 transition-colors cursor-pointer">
                <FileSpreadsheet className="w-5 h-5 mr-2" />
                Choose File
                <input
                  type="file"
                  className="hidden"
                  accept=".xlsx,.csv"
                  onChange={handleFileInput}
                />
              </label>

              <div className="mt-4 text-sm text-gray-500">
                Supported formats: Excel (.xlsx) and CSV files
              </div>
            </div>
          ) : (
            <div className="space-y-6">
              {/* File Preview */}
              <div className="bg-green-50 border border-green-200 rounded-lg p-6">
                <div className="flex items-center justify-between">
                  <div className="flex items-center space-x-3">
                    <CheckCircle className="w-8 h-8 text-green-600" />
                    <div>
                      <h4 className="font-semibold text-green-800">{file.name}</h4>
                      <p className="text-sm text-green-600">
                        {(file.size / 1024).toFixed(1)} KB • Ready to process
                      </p>
                    </div>
                  </div>
                  <button
                    onClick={resetUpload}
                    className="p-2 text-green-600 hover:bg-green-100 rounded-full transition-colors cursor-pointer"
                  >
                    <X className="w-5 h-5" />
                  </button>
                </div>
              </div>

              {/* Upload Button */}
              <div className="flex justify-center">
                {uploadStatus !== 'success' && (
                  <button
                  onClick={handleUpload}
                  disabled={uploading}
                  className="px-8 py-3 bg-[#6699ff] text-white font-semibold rounded-lg hover:bg-blue-700 disabled:opacity-50 disabled:cursor-not-allowed cursor-pointer transition-colors flex items-center space-x-2"
                >
                  {uploading ? (
                    <>
                      <div className="w-5 h-5 border-2 border-white/30 border-t-white rounded-full animate-spin"></div>
                      <span>Processing...</span>
                    </>
                  ) : (
                    <>
                      <Upload className="w-5 h-5" />
                      <span>Process File</span>
                    </>
                  )}
                </button>
                )}
                
              </div>
            </div>
          )}

          {/* Status Messages */}
          {uploadStatus === 'success' && (
            <div className="mt-6 bg-green-50 border flex justify-between border-green-200 rounded-lg p-4">
              <div className="flex items-center space-x-3">
                <CheckCircle className="w-6 h-6 text-green-600" />
                <div>
                  <h4 className="font-semibold text-green-800">File processed successfully!</h4>
                  
                </div>
              </div>
              <button className='px-4 py-2 bg-[#6699ff] text-white font-semibold rounded-lg hover:bg-blue-700 disabled:opacity-50 disabled:cursor-not-allowed cursor-pointer transition-colors'>Send Reminders</button>
            </div>
          )}

          {uploadStatus === 'error' && (
            <div className="mt-6 bg-red-50 border border-red-200 rounded-lg p-4">
              <div className="flex items-center space-x-3">
                <AlertCircle className="w-6 h-6 text-red-600" />
                <div>
                  <h4 className="font-semibold text-red-800">Invalid file format</h4>
                  <p className="text-sm text-red-700">Please upload a valid Excel (.xlsx) or CSV file.</p>
                </div>
              </div>
            </div>
          )}
        </div>

        {/* Instructions */}
        <div className="mt-12 bg-blue-50 rounded-xl p-8">
          <h4 className="text-lg font-semibold text-blue-900 mb-4">File Format Requirements</h4>
          <div className="grid md:grid-cols-2 gap-6 text-sm">
            <div>
              <h5 className="font-medium text-blue-800 mb-2">Required Columns:</h5>
              <ul className="space-y-1 text-blue-700">
                <li>• Student Name</li>
                <li>• Phone Number</li>
                <li>• Fee Balance</li>
              </ul>
            </div>
            <div>
              <h5 className="font-medium text-blue-800 mb-2">Example Format:</h5>
              <div className="bg-white rounded p-3 font-mono text-xs">
                Name, Phone, Balance<br/>
                John Doe, +234123456789, 50000<br/>
                Jane Smith, +234987654321, 0
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Home;