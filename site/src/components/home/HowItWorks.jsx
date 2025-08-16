import React from 'react';
import { ArrowRight, GitBranch, Settings, FolderOpen, Globe, Terminal } from 'lucide-react';

const steps = [
  {
    icon: Terminal,
    title: 'Repo URL or Current Repo',
    description: 'Pass any GitHub, GitLab, or Bitbucket URL',
    color: 'from-blue-500 to-cyan-500'
  },
  {
    icon: Settings,
    title: 'Parse & Resolve',
    description: 'Extract host/owner/repo and read ~/.gitr.yaml config',
    color: 'from-purple-500 to-pink-500'
  },
  {
    icon: FolderOpen,
    title: 'Compute Path',
    description: 'Generate deterministic directory structure',
    color: 'from-green-500 to-emerald-500'
  },
  {
    icon: GitBranch,
    title: 'Clone Smart',
    description: 'HTTPS via go-git or SSH via git(1) with auth',
    color: 'from-orange-500 to-red-500'
  },
  {
    icon: Globe,
    title: 'Navigate Web',
    description: 'Open any repo page with single commands',
    color: 'from-indigo-500 to-blue-500'
  }
];

export default function HowItWorks() {
  return (
    <div className="max-w-7xl mx-auto px-6">
      <div className="text-center mb-16">
        <h2 className="text-4xl md:text-5xl font-bold mb-6 bg-gradient-to-r from-white to-gray-300 bg-clip-text text-transparent">
          How it works
        </h2>
        <p className="text-xl text-gray-400 max-w-3xl mx-auto">
          One command, multiple actions. gitr handles the complexity so you can focus on code.
        </p>
      </div>

      {/* Flow Diagram */}
      <div className="relative mb-16">
        <div className="flex flex-col lg:flex-row items-center justify-between gap-8">
          {steps.map((step, index) => (
            <React.Fragment key={index}>
              <div className="flex flex-col items-center text-center max-w-xs">
                <div className={`w-16 h-16 rounded-2xl bg-gradient-to-r ${step.color} p-4 mb-4 shadow-lg`}>
                  <step.icon className="w-8 h-8 text-white" />
                </div>
                <h3 className="text-lg font-semibold text-white mb-2">
                  {step.title}
                </h3>
                <p className="text-gray-400 text-sm">
                  {step.description}
                </p>
              </div>
              {index < steps.length - 1 && (
                <ArrowRight className="hidden lg:block w-6 h-6 text-gray-600 flex-shrink-0" />
              )}
            </React.Fragment>
          ))}
        </div>
      </div>

      {/* Detailed Flow */}
      <div className="bg-gradient-to-r from-gray-900/50 to-gray-800/50 backdrop-blur-xl border border-gray-700 rounded-3xl p-8">
        <h3 className="text-2xl font-bold text-white mb-6 text-center">
          The complete flow
        </h3>
        <div className="space-y-6">
          <div className="flex items-start gap-4">
            <div className="w-8 h-8 rounded-full bg-blue-600 flex items-center justify-center text-white text-sm font-bold flex-shrink-0">
              1
            </div>
            <div>
              <h4 className="text-white font-semibold mb-1">Parse URL structure</h4>
              <p className="text-gray-400 text-sm">Extract provider, host, owner, and repository name from any valid git URL</p>
            </div>
          </div>
          <div className="flex items-start gap-4">
            <div className="w-8 h-8 rounded-full bg-purple-600 flex items-center justify-center text-white text-sm font-bold flex-shrink-0">
              2
            </div>
            <div>
              <h4 className="text-white font-semibold mb-1">Read configuration</h4>
              <p className="text-gray-400 text-sm">Load ~/.gitr.yaml for provider mappings, authentication, and path preferences</p>
            </div>
          </div>
          <div className="flex items-start gap-4">
            <div className="w-8 h-8 rounded-full bg-green-600 flex items-center justify-center text-white text-sm font-bold flex-shrink-0">
              3
            </div>
            <div>
              <h4 className="text-white font-semibold mb-1">Derive repository path</h4>
              <p className="text-gray-400 text-sm">Create deterministic local path: ~/scm/host.com/owner/repo</p>
            </div>
          </div>
          <div className="flex items-start gap-4">
            <div className="w-8 h-8 rounded-full bg-orange-600 flex items-center justify-center text-white text-sm font-bold flex-shrink-0">
              4
            </div>
            <div>
              <h4 className="text-white font-semibold mb-1">Plan and preview</h4>
              <p className="text-gray-400 text-sm">Use --dry flag to see exactly what gitr will do before making changes</p>
            </div>
          </div>
          <div className="flex items-start gap-4">
            <div className="w-8 h-8 rounded-full bg-red-600 flex items-center justify-center text-white text-sm font-bold flex-shrink-0">
              5
            </div>
            <div>
              <h4 className="text-white font-semibold mb-1">Clone with authentication</h4>
              <p className="text-gray-400 text-sm">HTTPS tokens or SSH keys automatically discovered and used</p>
            </div>
          </div>
          <div className="flex items-start gap-4">
            <div className="w-8 h-8 rounded-full bg-indigo-600 flex items-center justify-center text-white text-sm font-bold flex-shrink-0">
              6
            </div>
            <div>
              <h4 className="text-white font-semibold mb-1">Navigate and open</h4>
              <p className="text-gray-400 text-sm">Copy cd command to clipboard and open browser deep-links on demand</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}