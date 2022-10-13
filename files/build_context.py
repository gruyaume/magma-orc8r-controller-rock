#!/usr/bin/env python3

"""
Copyright 2020 The Magma Authors.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
"""

# build.py creates the build context for the orc8r Docker builds.
# It first creates a tmp directory, and then copies the cloud directories
# for all modules into it.

import glob
import os
import shutil
from collections import namedtuple
from typing import Iterable

HOST_BUILD_CTX = '/tmp/magma_orc8r_build'
HOST_MAGMA_ROOT = '../../../.'
IMAGE_MAGMA_ROOT = os.path.join('src', 'magma')


MODULES = [
    'orc8r',
    'lte',
    'feg',
    'cwf',
    'wifi',
    'fbinternal',
]

DEPLOYMENT_TO_MODULES = {
    'all': MODULES,
    'orc8r': ['orc8r'],
    'orc8r-f': ['orc8r', 'fbinternal'],
    'fwa': ['orc8r', 'lte'],
    'fwa-f': ['orc8r', 'lte', 'fbinternal'],
    'ffwa': ['orc8r', 'lte', 'feg'],
    'ffwa-f': ['orc8r', 'lte', 'feg', 'fbinternal'],
    'cwf': ['orc8r', 'lte', 'feg', 'cwf'],
    'cwf-f': ['orc8r', 'lte', 'feg', 'cwf', 'fbinternal'],
    'wifi': ['orc8r', 'wifi'],
    'wifi-f': ['orc8r', 'wifi', 'fbinternal'],
}

DEPLOYMENTS = DEPLOYMENT_TO_MODULES.keys()

MagmaModule = namedtuple('MagmaModule', ['name', 'host_path'])


def main() -> None:
    deployment = "all"
    mods = _get_modules(DEPLOYMENT_TO_MODULES[deployment])
    _create_build_context(mods)


def _get_modules(mods: Iterable[str]) -> Iterable[MagmaModule]:
    """
    Read the modules config file and return all modules specified.
    """
    modules = []
    for m in mods:
        abspath = os.path.abspath(os.path.join(HOST_MAGMA_ROOT, m))
        module = MagmaModule(name=m, host_path=abspath)
        modules.append(module)
    return modules


def _create_build_context(modules: Iterable[MagmaModule]) -> None:
    """ Clear out the build context from the previous run """
    shutil.rmtree(HOST_BUILD_CTX, ignore_errors=True)
    os.mkdir(HOST_BUILD_CTX)

    print("Creating build context in '%s'..." % HOST_BUILD_CTX)
    for m in modules:
        _copy_module(m)


def _copy_module(module: MagmaModule) -> None:
    """ Copy module directory into the build context  """
    build_ctx = _get_module_host_dst(module)

    def copy_to_ctx(d: str) -> None:
        shutil.copytree(
            os.path.join(module.host_path, d),
            os.path.join(build_ctx, d),
        )

    copy_to_ctx('cloud')

    # Orc8r module also has lib/ and gateway/
    if module.name == 'orc8r':
        copy_to_ctx('lib')
        copy_to_ctx('gateway')

    # Optionally copy cloud/configs/
    # Converts e.g. lte/cloud/configs/ to configs/lte/
    if os.path.isdir(os.path.join(module.host_path, 'cloud', 'configs')):
        shutil.copytree(
            os.path.join(module.host_path, 'cloud', 'configs'),
            os.path.join(HOST_BUILD_CTX, 'configs', module.name),
        )

    # Copy the go.mod file for caching the go downloads
    # Preserves relative paths between modules
    for f in glob.iglob(build_ctx + '/**/go.mod', recursive=True):
        gomod = f.replace(
            HOST_BUILD_CTX, os.path.join(HOST_BUILD_CTX, 'gomod'),
        )
        print(gomod)
        os.makedirs(os.path.dirname(gomod))
        shutil.copyfile(f, gomod)


def _get_module_host_dst(module: MagmaModule) -> str:
    """
    Given a path to a module on the host, return the intended destination
    in the build context.
    """
    return os.path.join(HOST_BUILD_CTX, IMAGE_MAGMA_ROOT, module.name)


if __name__ == '__main__':
    main()
