import React from 'react';
import { Terminal, Heart } from 'lucide-react';

export default function Footer() {
  const currentYear = new Date().getFullYear();

  return (
    <footer className="relative border-t border-gray-800 bg-black/50 backdrop-blur-xl">
      <div className="max-w-7xl mx-auto px-6 py-12">
        <div className="text-center">
          {/* Brand */}
          <div className="flex items-center justify-center gap-3 mb-6">
            <div className="w-8 h-8 bg-gradient-to-r from-blue-500 to-purple-500 rounded-lg flex items-center justify-center">
              <Terminal className="w-5 h-5 text-white" />
            </div>
            <span className="text-xl font-bold bg-gradient-to-r from-blue-400 to-purple-400 bg-clip-text text-transparent">
              gitr
            </span>
          </div>
          
          <p className="text-gray-400 text-sm mb-6 max-w-md mx-auto">
            Deterministic git clone paths and instant web deep-links for GitHub, GitLab, and Bitbucket.
          </p>

          {/* Copyright */}
          <div className="flex items-center justify-center gap-2 text-gray-500 text-sm">
            <span>© {currentYear} gitr. Made with</span>
            <Heart className="w-4 h-4 text-red-400 fill-current" />
            <span>for developers.</span>
          </div>
        </div>
      </div>
    </footer>
  );
}