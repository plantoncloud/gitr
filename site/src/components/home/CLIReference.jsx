import React from 'react';
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import { Card, CardContent, CardHeader } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { Terminal, Settings, FileCode } from 'lucide-react';

const CodeBlock = ({ children, title }) => (
  <div className="bg-black/50 border border-gray-800 rounded-xl overflow-hidden mb-4">
    {title && (
      <div className="px-4 py-2 bg-gray-900 border-b border-gray-800">
        <span className="text-gray-400 text-sm">{title}</span>
      </div>
    )}
    <pre className="p-4 overflow-x-auto">
      <code className="text-gray-300 font-mono text-sm whitespace-pre-wrap">
        {children}
      </code>
    </pre>
  </div>
);

const commands = [
  {
    command: 'gitr clone <url>',
    description: 'Clone repository to deterministic path',
    example: 'gitr clone https://github.com/owner/repo',
    flags: ['--create-dir', '--token', '--dry']
  },
  {
    command: 'gitr web',
    description: 'Open repository home page',
    example: 'gitr web',
    flags: ['--dry']
  },
  {
    command: 'gitr prs',
    description: 'Open pull requests / merge requests',
    example: 'gitr prs',
    flags: ['--dry']
  },
  {
    command: 'gitr pipelines',
    description: 'Open GitHub Actions / GitLab Pipelines',
    example: 'gitr pipelines',
    flags: ['--dry']
  },
  {
    command: 'gitr issues',
    description: 'Open issues page',
    example: 'gitr issues',
    flags: ['--dry']
  },
  {
    command: 'gitr rem',
    description: 'Open current branch in web UI',
    example: 'gitr rem',
    flags: ['--dry']
  }
];

const flags = [
  {
    flag: '--dry',
    description: 'Dry run - preview what gitr will do without making changes',
    scope: 'Global'
  },
  {
    flag: '--debug',
    description: 'Set log level to debug for detailed output',
    scope: 'Global'
  },
  {
    flag: '--create-dir',
    description: 'Create folders mirroring the repo path on the host',
    scope: 'clone/path'
  },
  {
    flag: '--token <string>',
    description: 'HTTPS personal access token for authentication',
    scope: 'clone'
  }
];

export default function CLIReference() {
  return (
    <div className="max-w-7xl mx-auto px-6">
      <div className="text-center mb-16">
        <h2 className="text-4xl md:text-5xl font-bold mb-6 bg-gradient-to-r from-white to-gray-300 bg-clip-text text-transparent">
          CLI Reference
        </h2>
        <p className="text-xl text-gray-400 max-w-3xl mx-auto">
          Complete command reference with examples and configuration options
        </p>
      </div>

      <Tabs defaultValue="commands" className="max-w-5xl mx-auto">
        <TabsList className="bg-gray-900 border border-gray-700 grid grid-cols-3 w-full">
          <TabsTrigger value="commands" className="flex items-center gap-2">
            <Terminal className="w-4 h-4" />
            Commands
          </TabsTrigger>
          <TabsTrigger value="flags" className="flex items-center gap-2">
            <Settings className="w-4 h-4" />
            Flags
          </TabsTrigger>
          <TabsTrigger value="config" className="flex items-center gap-2">
            <FileCode className="w-4 h-4" />
            Config
          </TabsTrigger>
        </TabsList>

        <TabsContent value="commands" className="mt-8">
          <div className="space-y-6">
            <div className="grid md:grid-cols-2 gap-6">
              {commands.map((cmd, index) => (
                <Card key={index} className="bg-gray-900/50 border-gray-700">
                  <CardHeader className="pb-3">
                    <div className="flex items-center justify-between">
                      <code className="text-blue-400 font-mono text-sm">
                        {cmd.command}
                      </code>
                    </div>
                    <p className="text-gray-400 text-sm">{cmd.description}</p>
                  </CardHeader>
                  <CardContent>
                    <CodeBlock>
                      $ {cmd.example}
                    </CodeBlock>
                    <div className="flex flex-wrap gap-2">
                      {cmd.flags.map((flag) => (
                        <Badge key={flag} variant="secondary" className="bg-gray-800 text-gray-300">
                          {flag}
                        </Badge>
                      ))}
                    </div>
                  </CardContent>
                </Card>
              ))}
            </div>
          </div>
        </TabsContent>

        <TabsContent value="flags" className="mt-8">
          <div className="space-y-4">
            {flags.map((flag, index) => (
              <Card key={index} className="bg-gray-900/50 border-gray-700">
                <CardContent className="p-6">
                  <div className="flex items-start justify-between mb-2">
                    <code className="text-green-400 font-mono">
                      {flag.flag}
                    </code>
                    <Badge variant="secondary" className="bg-gray-800 text-gray-300">
                      {flag.scope}
                    </Badge>
                  </div>
                  <p className="text-gray-400">{flag.description}</p>
                </CardContent>
              </Card>
            ))}
          </div>
        </TabsContent>

        <TabsContent value="config" className="mt-8">
          <div className="space-y-6">
            <Card className="bg-gray-900/50 border-gray-700">
              <CardHeader>
                <h3 className="text-xl font-semibold text-white flex items-center gap-2">
                  <Settings className="w-5 h-5" />
                  Configuration Management
                </h3>
              </CardHeader>
              <CardContent className="space-y-4">
                <CodeBlock title="Initialize config">
{`$ gitr config init  # creates ~/.gitr.yaml with defaults if missing`}
                </CodeBlock>
                <CodeBlock title="View current config">
{`$ gitr config show  # prints effective configuration`}
                </CodeBlock>
                <CodeBlock title="Edit config">
{`$ gitr config edit  # opens ~/.gitr.yaml in VS Code`}
                </CodeBlock>
              </CardContent>
            </Card>

            <Card className="bg-gray-900/50 border-gray-700">
              <CardHeader>
                <h3 className="text-xl font-semibold text-white">Example gitr.yaml</h3>
              </CardHeader>
              <CardContent>
                <CodeBlock>
{`# ~/.gitr.yaml
scm:
  homeDir: ~/scm  # base directory for all repos

clone:
  alwaysCreateDir: true
  includeHostForCreateDir: true
  defaultBranch: main

auth:
  httpsTokenFile: ~/.personal_access_tokens
  sshConfigFile: ~/.ssh/config`}
                </CodeBlock>
              </CardContent>
            </Card>
          </div>
        </TabsContent>
      </Tabs>
    </div>
  );
}