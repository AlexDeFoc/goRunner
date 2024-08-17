#include "env.iss"

[Code]
procedure CurStepChanged(CurStep: TSetupStep);
begin
    if CurStep = ssPostInstall 
     then EnvAddPath(ExpandConstant('{app}'));
end;

procedure CurUninstallStepChanged(CurUninstallStep: TUninstallStep);
begin
    if CurUninstallStep = usPostUninstall
    then EnvRemovePath(ExpandConstant('{app}'));
end;

[Setup]
AppName=goRunner
AppVersion=1.0
DefaultDirName={commonpf64}\goRunner
DefaultGroupName=goRunner
OutputDir=../installer
OutputBaseFilename=win_amd64
Compression=lzma2/ultra64
SolidCompression=yes
WizardStyle=modern
ChangesEnvironment=yes

[Files]
Source: "C:/Users/Alefan/Documents/work/goRunner/r.exe"; DestDir: "{app}"; Flags: ignoreversion
Source: "C:/Users/Alefan/Documents/work/goRunner/config.json"; DestDir: "{app}"; Flags: ignoreversion
Source: "C:/Users/Alefan/Documents/work/goRunner/app.json"; DestDir: "{app}"; Flags: ignoreversion
Source: "C:/Users/Alefan/Documents/work/goRunner/alias.json"; DestDir: "{app}"; Flags: ignoreversion

[Icons]
Name: "{group}\goRunner"; Filename: "{app}\r.exe"; WorkingDir: "{app}"; IconFilename: "{app}\r.exe"; IconIndex: 0
Name: "{commonstartup}\goRunner"; Filename: "{app}\r.exe"; WorkingDir: "{app}"; IconFilename: "{app}\r.exe"; IconIndex: 0
