/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API Playlist Vidoe
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntScholarshipRequest,
    EntScholarshipRequestFromJSON,
    EntScholarshipRequestFromJSONTyped,
    EntScholarshipRequestToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntSemesterEdges
 */
export interface EntSemesterEdges {
    /**
     * ScholarshipRequest holds the value of the ScholarshipRequest edge.
     * @type {Array<EntScholarshipRequest>}
     * @memberof EntSemesterEdges
     */
    scholarshipRequest?: Array<EntScholarshipRequest>;
}

export function EntSemesterEdgesFromJSON(json: any): EntSemesterEdges {
    return EntSemesterEdgesFromJSONTyped(json, false);
}

export function EntSemesterEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntSemesterEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'scholarshipRequest': !exists(json, 'scholarshipRequest') ? undefined : ((json['scholarshipRequest'] as Array<any>).map(EntScholarshipRequestFromJSON)),
    };
}

export function EntSemesterEdgesToJSON(value?: EntSemesterEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'scholarshipRequest': value.scholarshipRequest === undefined ? undefined : ((value.scholarshipRequest as Array<any>).map(EntScholarshipRequestToJSON)),
    };
}


